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
import { useNavigate } from "react-router-dom";
import { useEffect } from 'react';
import axios from 'axios';

function Home({active_item}) {
  
  const navigate = useNavigate();
  var x = document.cookie
  const dict = {}
  let chatsItems = [
    {id:1, photo:image1, time:"12 sec", name:"Parsa sabaghi", message:"kheili jalebe!", isOnline:true, newMessages:3},
    {id:2, photo:image2, time:"2 hours", name:"Hamid rezaei", message:"ahaaaa", isOnline:false, newMessages:18},
    {id:3, photo:null, time:"1 year", name:"no one", message:"ee", isOnline:false, newMessages:5}
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
  

  // for(let i of document.cookie.split(","))
  //   dict[i.split("=")[0]] = i.split("=")[1]
  // if(dict["token"] == undefined){
  //   navigate("/login");
  // }
  //   else{
  //     await axios.get("http://localhost:8080/api/chats", {headers: {
  //       Authorization: `Bearer ${dict["token"]}`
  //     },}).then(async (response) => {
	// 			if(response.status == 200){
  //         console.log("8888")
  //         console.log(response.data.Chats)
	// 				if(response.data.Chats != null){
  //           for(let chat of response.data.Chats){
  //             chatsItems.push({id: chat.ID, photo:image1, time:"12 sec", name:"parsa sabaghi", message: "salam", isOnline:false, newMessages:10})
  //           }
  //           console.log("pppp")
  //           console.log(chatsItems)
  //             await axios.get(`http://localhost:8080/api/chats/${chatsItems[0].id}`, {headers: {
  //               Authorization: `Bearer ${dict["token"]}`
  //             },}).then(async (response) => {
  //               if((await response).status == 200){
  //                 if(response.data.messages != null){
  //                   for(let message of response.data.messages)
  //                     messages.push({id: message.ID, photo:image1, time:"12 sec", text:"salam!!!! khooobi?", isOnline:false, isMine:(message.sender_id == dict["id"])})
  //                 }
  //                 console.log("iiii")
  //           console.log(messages)
  //               }
  //               }); 
  //         }
	// 			}
	// 	  	}); 
  //   }
  //   console.log("yyyy")
  // console.log(messages)
  
  let menuItems=[
    {id:"profileButton", image:profileImage, name:"profile"},
    {id:"chatsButton", image:chatsImage, name:"chats"}
  ]

  return (
    <>
    <div className={style.container}>
      <div className={style.row}>
        <SideMenu active_item={active_item} items={menuItems}></SideMenu>
        <ChatsSection items={chatsItems} activeItem={1}></ChatsSection>
        <section className={style.chat}>
        <ChatHeader item={chatsItems[0]?chatsItems[0]:{id:3, photo:null, time:"1 year", name:"no one", message:"ee", isOnline:false, newMessages:5}}></ChatHeader>
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
