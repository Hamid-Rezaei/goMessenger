import style from '../pages/HomePage.module.css';
import sendImage from '../statics/send.png'

function SendField() {
  return (
    <>
	<div className={style.footer_chat}>
          <input type="text" className={style.write_message} placeholder="Type your message here"></input>
          <img src={sendImage} alt={"send"} className={style.send_button} id="sendButton"/>
        </div>
    </>
  )
}

export default SendField;