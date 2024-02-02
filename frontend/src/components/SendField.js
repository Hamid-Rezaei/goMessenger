import style from '../pages/HomePage.module.css';
import sendImage from '../statics/send.png'

function SendField({func, setText, text}) {
  return (
    <>
	<div className={style.footer_chat}>
          <input type="text" className={style.write_message} placeholder="Type your message here" onChange = {(e) => setText(e.target.value)}></input>
          <img src={sendImage} alt={"send"} className={style.send_button} onClick={func} id="sendButton"/>
        </div>
    </>
  )
}

export default SendField;