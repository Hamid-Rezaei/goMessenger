import style from './Login.module.css';
import image from '../statics/icon.png'
import FormButtonButton from '../components/FormButtonButton';
import usernameImage from '../statics/username.png'
import passwordImage from '../statics/password.png'
import Field from '../components/Field';

function Login() {
  return (
    <>
<div className={style.container}>
	<div className={style.screen}>
		<div className={style.screen__content}>
      		<div className={style.logo_container}>
            	<img src={image} alt="Avatar" className={style.icon}/>
				<div className={style.title}>Login</div>
      		</div>
			<form className={style.login}>
				<Field placeholder={"User name"} type={"text"} image={usernameImage}  id={"userName"}></Field>
				<Field placeholder={"Password"} type={"Password"} image={passwordImage}  id={"password"}></Field>
        		<div className={style.buttons_container}>
					<FormButtonButton className={style.login__submit} textClassName={style.button__text} text={'Log In'}></FormButtonButton>
					<FormButtonButton className={style.signup__button} textClassName={style.button__text} text={'Sign Up'}></FormButtonButton>	
        		</div>
			</form>
		</div>	
	</div>
</div>
    </>
  )
}

export default Login;
