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
import { useEffect, useState } from 'react';
import axios from 'axios';

function Home({active_item}) {
  
  const navigate = useNavigate();

  const [chatsItems, setchatsItems] = useState([]);
  const [messages, setMessagesItems] = useState([]);
  var x = document.cookie
  const dict = {}
  let chatsItems1 = [
    // {id:1, photo:image1, time:"12 sec", name:"Parsa sabaghi", message:"kheili jalebe!", isOnline:true, newMessages:3},
    // {id:2, photo:image2, time:"2 hours", name:"Hamid rezaei", message:"ahaaaa", isOnline:false, newMessages:18},
    // {id:3, photo:null, time:"1 year", name:"no one", message:"ee", isOnline:false, newMessages:5}
  ]
  let messages1=[
    // {photo:image1, iOnline:false, text:"hello", time:"14:34", isMine:false},
    // {photo:image1, iOnline:false,text:"how\n are\n you?\nha\n ha?", time:"14:34", isMine:false},
    // {photo:image3, iOnline:false,text:"hello", time:"14:34", isMine:true},
    // {photo:image1, iOnline:false,text:"hello", time:"14:34", isMine:false},
    // {photo:image3, iOnline:false,text:"hello", time:"14:34", isMine:true},
    // {photo:image1, iOnline:false,text:"hello", time:"14:34", isMine:false},
    // {photo:image3, iOnline:false,text:"hello", time:"14:34", isMine:true},
    // {photo:image1, iOnline:false,text:"hello", time:"14:34", isMine:false}
  ]


  const temp = async ()=>{
    for(let i of document.cookie.split(","))
      dict[i.split("=")[0]] = i.split("=")[1]
    if(dict["token"] == undefined){
      navigate("/login");
    }
    else{
      let flag = true
      await axios.get("http://localhost:8080/api/chats", {headers: {
        Authorization: `Bearer ${dict["token"]}`
      },}).catch(function (error) {
				if (error.response) {
				  // The request was made and the server responded with a status code
				  // that falls out of the range of 2xx
				  console.log(error.response.data);
				  console.log(error.response.status);
				  console.log(error.response.headers);
				} else if (error.request) {
				  // The request was made but no response was received
				  // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
				  // http.ClientRequest in node.js
				  console.log(error.request);
				} else {
				  // Something happened in setting up the request that triggered an Error
				  console.log('Error', error.message);
				}
				console.log(error.config);
				flag = false
			  }).then(async function (response) {
				if(flag){
          if(response.status == 200){
            console.log("8888")
            console.log(response.data.Chats)
            if(response.data.Chats != null){
              for(let chat of response.data.Chats){
                chatsItems1.push({id: chat.ID, photo:image1, time:"12 sec", name:"parsa sabaghi", message: "salam", isOnline:false, newMessages:10})
              }
              setchatsItems(chatsItems1)
              console.log("pppp")
              console.log(chatsItems)
                await axios.get(`http://localhost:8080/api/chats/${chatsItems1[0].id}`, {headers: {
                  Authorization: `Bearer ${dict["token"]}`
                },}).catch(function (error) {
                  if (error.response) {
                    // The request was made and the server responded with a status code
                    // that falls out of the range of 2xx
                    console.log(error.response.data);
                    console.log(error.response.status);
                    console.log(error.response.headers);
                  } else if (error.request) {
                    // The request was made but no response was received
                    // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
                    // http.ClientRequest in node.js
                    console.log(error.request);
                  } else {
                    // Something happened in setting up the request that triggered an Error
                    console.log('Error', error.message);
                  }
                  console.log(error.config);
                  flag = false
                  }).then(function (response) {
                  if(flag){
                    if(response.status == 200){
                      if(response.data.messages != null){
                        for(let message of response.data.messages)
                          messages1.push({id: message.ID, photo:message.sender_id == dict["id"]?image2:image1, time:"12 sec", text:message.content, isOnline:false, isMine:(message.sender_id == dict["id"])})
                      }
                      setMessagesItems(messages1)
                      console.log(messages)
                    }
                  }
                }) 
            }
          }
			  }
			}) 
    }
  }
  let fl = true
  useEffect(()=>{if(fl){temp(); fl=false}}, [])

  const [sendText, setSendText] = useState("");
  
  let menuItems=[
    {id:"profileButton", image:profileImage, name:"profile"},
    {id:"chatsButton", image:chatsImage, name:"chats"}
  ]

  async function sendNewMessage(){
    if(sendText != ""){
      for(let i of document.cookie.split(","))
      dict[i.split("=")[0]] = i.split("=")[1]
    if(dict["token"] == undefined){
      navigate("/login");
    }
      let flag = true
      await axios.post("http://localhost:8080/api/chats/" + chatsItems[0].id + "/messages", {content:sendText}, {headers: {
        Authorization: `Bearer ${dict["token"]}`
      },}).catch(function (error) {
			if (error.response) {
			  // The request was made and the server responded with a status code
			  // that falls out of the range of 2xx
			  console.log(error.response.data);
			  console.log(error.response.status);
			  console.log(error.response.headers);
			} else if (error.request) {
			  // The request was made but no response was received
			  // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
			  // http.ClientRequest in node.js
			  console.log(error.request);
			} else {
			  // Something happened in setting up the request that triggered an Error
			  console.log('Error', error.message);
			}
			console.log(error.config);
			flag = false
		    }).then(async function (response) {
		  	if(flag){
          if(response.status == 200){
            await axios.get(`http://localhost:8080/api/chats/${chatsItems[0].id}`, {headers: {
                  Authorization: `Bearer ${dict["token"]}`
                },}).catch(function (error) {
                  if (error.response) {
                    // The request was made and the server responded with a status code
                    // that falls out of the range of 2xx
                    console.log(error.response.data);
                    console.log(error.response.status);
                    console.log(error.response.headers);
                  } else if (error.request) {
                    // The request was made but no response was received
                    // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
                    // http.ClientRequest in node.js
                    console.log(error.request);
                  } else {
                    // Something happened in setting up the request that triggered an Error
                    console.log('Error', error.message);
                  }
                  console.log(error.config);
                  flag = false
                  }).then(function (response) {
                  if(flag){
                    if(response.status == 200){
                      if(response.data.messages != null){
                        for(let message of response.data.messages)
                          messages1.push({id: message.ID, photo:message.sender_id == dict["id"]?image2:image1, time:"12 sec", text:message.content, isOnline:false, isMine:(message.sender_id == dict["id"])})
                      }
                      setMessagesItems(messages1)
                      console.log(messages)
                    }
                  }
                }) 
          }
		    }
		  }) 
  }
  }
  

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
        <SendField func={sendNewMessage} setText={setSendText} text={sendText}></SendField>
        </section>
      </div>
    </div>
    </>
  )
}

export default Home;
