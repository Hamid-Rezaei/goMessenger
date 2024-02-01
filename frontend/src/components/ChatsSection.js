import style from '../pages/HomePage.module.css';
import SearchBar from './SearchBar';
import profileImage from '../statics/profile.png'


function ChatsSection({items, activeItem}) {
    return (
      <>
        <section className={style.discussions}>
            <SearchBar></SearchBar>
            {items.map(item=>
                <div className={item.id==activeItem ? style.message_active: style.discussion}>
                    <div className={style.photo}>
                      <img src={item.photo? item.photo: profileImage} alt={""} className={item.isOnline? style.profile_pic_online:style.profile_pic_ofline}/>
                    </div>
                    <div className={style.desc_contact}>
                        <p className={style.name}>{item.name}</p>
                        <p className={style.message}>{item.message}</p>
                    </div>
                    <div className={style.newMessages}>{item.newMessages}</div>
                    <div className={style.timer}>{item.time}</div>
                </div>
            )}
        </section>
      </>
    )
  }
  
  export default ChatsSection;
  