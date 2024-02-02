import style from './Login.module.css';
import image from '../statics/icon.png'
import FormButtonButton from '../components/FormButtonButton';
import usernameImage from '../statics/username.png'
import passwordImage from '../statics/password.png'
import Field from '../components/Field';
import { useState } from 'react';
import SignUp from './SignUp';
import React from 'react';
import ReactDOM from 'react-dom/client';
import axios from 'axios';
import Home from './HomePage';
import { useNavigate } from "react-router-dom";

function Login() {
	const navigate = useNavigate();
	const [userName, setUserName] = useState('');
	const [password, setPassword] = useState('');

	const handleSubmit = (e) => {
        e.preventDefault();
        axios.post("http://localhost:8080/api/users/login", {
			username: userName,
			password:password,
			}).then((response) => {
				console.log(response.data)
				if(response.status == 200){
					document.cookie = `token=${response.data.token},username=${response.data.username},id=${response.data.Id}`
					navigate("/");
				}
				
		  	}); 

    }

	const goToSignUp = (e) => {
		navigate("/signup");
	}

  return (
    <>
<div className={style.container}>
	<div className={style.screen}>
		<div className={style.screen__content}>
      		<div className={style.logo_container}>
            	<img src={image} alt="Avatar" className={style.icon}/>
				<div className={style.title}>Login</div>
      		</div>
			<form className={style.login} onSubmit = {handleSubmit}>
				<Field placeholder={"User name"} type={"text"} image={usernameImage}  id={"userName"} value={userName} setValue={setUserName}></Field>
				<Field placeholder={"Password"} type={"Password"} image={passwordImage}  id={"password"} value={password} setValue={setPassword}></Field>
        		<div className={style.buttons_container}>
					<FormButtonButton className={style.login__submit} textClassName={style.button__text} text={'Log In'} id={"submit"}></FormButtonButton>
					<FormButtonButton className={style.signup__button} textClassName={style.button__text} text={'Sign Up'} id={"signUp"} onClick={goToSignUp}></FormButtonButton>	
        		</div>
			</form>
		</div>	
	</div>
</div>
    </>
  )
}

export default Login;
