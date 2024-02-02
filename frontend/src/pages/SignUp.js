import style from './Login.module.css';
import image from '../statics/icon.png'
import FormButtonButton from '../components/FormButtonButton';
import usernameImage from '../statics/username.png'
import passwordImage from '../statics/password.png'
import phoneImage from '../statics/phone.png'
import nameImage from '../statics/name.png'
import Field from '../components/Field';
import axios from 'axios';
import { useState } from 'react';
import React from 'react';
import ReactDOM from 'react-dom/client';
import Login from './Login'
import { useNavigate } from "react-router-dom";

function SignUp() {
	const navigate = useNavigate();
	const [userName, setUserName] = useState('');
	const [userFName, setFName] = useState('');
	const [userLName, setLName] = useState('');
	const [bio, setBio] = useState('');
	const [phoneNumber, setPhoneNumber] = useState('');
	const [password, setPassword] = useState('');
	
    
    const handleSubmit = (e) => {
    
        e.preventDefault();
        axios.post("http://localhost:8080/api/users/register", {
			firstname: userFName,
			lastname: userLName,
			username: userName,
			password:password,
			phone: phoneNumber,
			bio: bio
			}).then((response) => {
				if(response.status == 201){
					goToLogin()
				}
				else{
					return "z"
				}
		  	}); 

    }

	const goToLogin = (e) => {
		navigate("/login");
	}

  return (
    <>
<div className={style.container}>
	<div className={style.screen}>
		<div className={style.screen__content}>
		<div className={style.logo_container}>
            	<img src={image} alt="Avatar" className={style.icon}/>
				<div className={style.title}>SignUp</div>
      		</div>
			<form className={style.login} onSubmit = {handleSubmit}>
				<div className={style.flex_horizontal}>
					<Field placeholder={"First name"} type={"text"} image={nameImage} id={"firstName"} value={userFName} setValue={setFName}></Field>
					<Field placeholder={"Last name"} type={"text"} image={nameImage} id={"lastName"} value={userLName} setValue={setLName}> </Field>
        		</div>
				<Field placeholder={"Phone number"} type={"text"} image={phoneImage} id={"phoneNumber"} value={phoneNumber} setValue={setPhoneNumber}></Field>
				<Field placeholder={"Bio"} type={"text"} image={nameImage} id={"bio"} value={bio} setValue={setBio}></Field>
				<Field placeholder={"User name"} type={"text"} image={usernameImage}  id={"userName"} value={userName} setValue={setUserName}></Field>
				<Field placeholder={"Password"} type={"Password"} image={passwordImage}  id={"password"} value={password} setValue={setPassword}></Field>
        		<div className={style.buttons_container}>
					<FormButtonButton className={style.login__submit}  textClassName={style.button__text} text={'SignUp'} id={"submit"}></FormButtonButton>
					<FormButtonButton className={style.login__button} textClassName={style.button__text} text={'Login'} id={"login"} onClick={goToLogin}></FormButtonButton>	
        		</div>
			</form>
		</div>	
	</div>
</div>

<script>
	
</script>
    </>
  )
}

export default SignUp;
