
import style from '../pages/Login.module.css';
function Field({type, placeholder, image}) {
    return (
      <>
        <div className={style.login__field}>
		    <div className={style.flex_horizontal}>
		    	<img src={image} alt="Avatar" className={style.field_icon}/>
		    	<input type={type} className={style.login__input} placeholder={placeholder}/>
		    </div>
		</div>
      </>
    )
  }
  
  export default Field;
  