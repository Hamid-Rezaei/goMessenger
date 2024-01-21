package model

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"image"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	Phone     string `json:"phone" gorm:"unique"`
	Image     string `json:"image"`
	Bio       string `json:"bio"`
}

func (u *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}

func generateUniqueFilename(originalFilename string) string {
	year := time.Now().Year()
	return fmt.Sprintf("%d_user_image_%s", year, originalFilename)
}

func (u *User) UploadImage(file *multipart.FileHeader) (string, error) {

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {

		}
	}(src)

	filename := generateUniqueFilename(file.Filename)
	dest := filepath.Join("upload", filename)

	dst, err := os.Create(dest)
	if err != nil {
		return "", err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {

		}
	}(dst)

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	return dest, nil
}

func (u *User) RetrieveImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	profile, _, err := image.Decode(f)
	return profile, err
}
