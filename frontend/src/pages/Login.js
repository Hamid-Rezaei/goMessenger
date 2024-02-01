import style from './Login.module.css';
import image from '../statics/icon.png'
import Button from '../components/Button';
import usernameImage from '../statics/username.png'
import passwordImage from '../statics/password.png'
import Field from '../components/Field';

function Login() {
  return (
    <>
<div className={style.container}>
	<div className={style.screen}>
		<div className={style.screen__content}>
      		<div className={style.flex_vertical}>
        		<div className='empty'></div>
        		<div className={style.flex_horizontal}>
          			<div className='empty'></div>
            		<img src={image} alt="Avatar" className={style.icon}/>
          			<div className='empty'></div>
        		</div>
        		<div className='empty'><span className={style.title}>Login</span></div>
      		</div>
			<form className={style.login}>
				<Field placeholder={"User name"} type={"text"} image={usernameImage}></Field>
				<Field placeholder={"Password"} type={"Password"} image={passwordImage}></Field>
        		<div className={style.buttons_container}>
					<Button className={style.login__submit} textClassName={style.button__text} text={'Log In'}></Button>
					<Button className={style.signup__button} textClassName={style.button__text} text={'Sign Up'}></Button>	
        		</div>
			</form>
		</div>	
	</div>
</div>
    </>
  )
}

export default Login;
