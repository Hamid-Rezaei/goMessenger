import style from './HomePage.module.css';
import profileImage from '../statics/profile.png'
import chatsImage from '../statics/chats.png'
import image1 from '../statics/1.jpg'
import image2 from '../statics/2.jpg'
import image3 from '../statics/3.jpg'
import SideMenu from '../components/SideMenu';
import ChatsSection from '../components/ChatsSection';
import ChatHeader from '../components/ChatHeader';
import Message from '../components/Message';
import SendField from '../components/SendField';

function Home({active_item}) {
  let chatsItems = [
    {id:1, photo:image1, time:"12 sec", name:"Parsa sabaghi", message:"kheili jalebe!", isOnline:true, newMessages:3},
    {id:2, photo:image2, time:"2 hours", name:"Hamid rezaei", message:"ahaaaa", isOnline:false, newMessages:18},
    {id:3, photo:null, time:"1 year", name:"no one", message:"ee", isOnline:false, newMessages:5}
  ]
  let menuItems=[{id:"profileButton", image:profileImage, name:"profile"},
   {id:"chatsButton", image:chatsImage, name:"chats"}
  ]
  let messages=[
    {photo:image1, iOnline:false, text:"hello", time:"14:34", isMine:false},
    {photo:image1, iOnline:false,text:"how\n are\n you?\nha\n ha?", time:"14:34", isMine:false},
    {photo:image3, iOnline:false,text:"hello", time:"14:34", isMine:true},
    {photo:image1, iOnline:false,text:"hello", time:"14:34", isMine:false},
    {photo:image3, iOnline:false,text:"hello", time:"14:34", isMine:true},
    {photo:image1, iOnline:false,text:"hello", time:"14:34", isMine:false},
    {photo:image3, iOnline:false,text:"hello", time:"14:34", isMine:true},
    {photo:image1, iOnline:false,text:"hello", time:"14:34", isMine:false}
  ]
  return (
    <>
    <div className={style.container}>
      <div className={style.row}>
        <SideMenu active_item={active_item} items={menuItems}></SideMenu>
        <ChatsSection items={chatsItems} activeItem={1}></ChatsSection>
        <section className={style.chat}>
        <ChatHeader item={chatsItems[0]}></ChatHeader>
        <div className={style.messages_chat}>
          {messages.map(item => <Message item={item}></Message>)}
        </div>
        <SendField></SendField>
        </section>
      </div>
    </div>
    </>
  )
}

export default Home;
