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
		var flag = true
        axios.post("http://localhost:8080/api/users/login", {
			username: userName,
			password:password,
			}).catch(function (error) {
				if (error.response) {
				  // The request was made and the server responded with a status code
				  // that falls out of the range of 2xx
				  console.log(error.response.data);
				  console.log(error.response.status);
				  console.log(error.response.headers);
				} else if (error.request) {
				  // The request was made but no response was received
				  // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
				  // http.ClientRequest in node.js
				  console.log(error.request);
				} else {
				  // Something happened in setting up the request that triggered an Error
				  console.log('Error', error.message);
				}
				console.log(error.config);
				document.getElementById("error").innerText = "invalid creditionals"
				flag = false
			  }).then(function (response) {
				if(flag){
				if(response.status == 200){
				document.cookie = `token=${response.data.token},username=${response.data.username},id=${response.data.Id}`
				navigate("/");
				}
			  }
			})
			
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
				<p id="error" style={{color : "red"}}></p>
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
