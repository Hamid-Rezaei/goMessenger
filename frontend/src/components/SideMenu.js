import style from '../pages/HomePage.module.css';
import MenuItem from "./MenuItem"

function SideMenu({active_item, items}) {
  return (
    <>
	<nav className={style.menu}>
        <ul className={style.items}>
            {items.map((item => 
                <li className={(active_item == item.name) ? style.item_active:style.item}>
                    <MenuItem id={item.id} image={item.image} key={item.id}></MenuItem>
                </li>)
            )}
        </ul>
      </nav>
    </>
  )
}

export default SideMenu;
