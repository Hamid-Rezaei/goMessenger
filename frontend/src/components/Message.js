import style from '../pages/HomePage.module.css';
import SearchBar from './SearchBar';
import profileImage from '../statics/profile.png'


function Message({item}) {
    return (
      <>
        <div className={item.isMine? style.message_in_chat_mine : style.message_in_chat_not_mine}>
            <div className={style.photo}>
                <img src={item.photo? item.photo: profileImage} alt={""} className={item.isOnline || item.isMine? style.profile_pic_online:style.profile_pic_ofline}/>
            </div>
            <pre className={style.text}> {item.text} </pre>
            <p className={style.time}> {item.time}</p>
          </div>
      </>
    )
  }
  
  export default Message;
  