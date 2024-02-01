import style from './Login.module.css';
import image from '../statics/icon.png'
import FormButtonButton from '../components/FormButtonButton';
import usernameImage from '../statics/username.png'
import passwordImage from '../statics/password.png'
import phoneImage from '../statics/phone.png'
import nameImage from '../statics/name.png'
import Field from '../components/Field';

function SignUp() {
  return (
    <>
<div className={style.container}>
	<div className={style.screen}>
		<div className={style.screen__content}>
		<div className={style.logo_container}>
            	<img src={image} alt="Avatar" className={style.icon}/>
				<div className={style.title}>SignUp</div>
      		</div>
			<form className={style.login}>
				<div className={style.flex_horizontal}>
					<Field placeholder={"First name"} type={"text"} image={nameImage} id={"firstName"}></Field>
					<Field placeholder={"Last name"} type={"text"} image={nameImage} id={"lastName"}></Field>
        		</div>
				<Field placeholder={"Phone number"} type={"text"} image={phoneImage} id={"phoneNumber"}></Field>
				<Field placeholder={"User name"} type={"text"} image={usernameImage}  id={"userName"}></Field>
				<Field placeholder={"Password"} type={"Password"} image={passwordImage}  id={"password"}></Field>
        		<div className={style.buttons_container}>
					<FormButtonButton className={style.login__submit}  textClassName={style.button__text} text={'Log In'} id={"submit"}></FormButtonButton>
					<FormButtonButton className={style.login__button} textClassName={style.button__text} text={'Login'} id={"login"}></FormButtonButton>	
        		</div>
			</form>
		</div>	
	</div>
</div>
    </>
  )
}

export default SignUp;
