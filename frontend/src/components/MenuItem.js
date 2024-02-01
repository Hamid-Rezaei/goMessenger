import style from '../pages/HomePage.module.css';

function MenuItem({image, id}) {
  return (
    <>
	    <div className={style.icon_container}>
          <img src={image} alt={id} className={style.menu_icon} id={id}/>
      </div>
    </>
  )
}

export default MenuItem;
