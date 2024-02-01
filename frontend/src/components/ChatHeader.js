import style from '../pages/HomePage.module.css';
import profileImage from '../statics/profile.png'

function ChatHeader({item}) {
  return (
    <>
	<div className={style.header_chat}>
        <img src={item.photo? item.photo: profileImage} alt={""} className={style.chat_profile_pic}/>
        <p className={style.name}>{item.name}</p>
    </div>
    </>
  )
}

export default ChatHeader;