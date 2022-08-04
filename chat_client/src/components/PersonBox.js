const PersonBox = (props) => {
    let msgsLen = (props.chatroom.messages).length
    let lastMessage = props.chatroom.messages[msgsLen - 1].text
    let lastMessageDateTime = props.chatroom.messages[msgsLen - 1].send_date
    let date = lastMessageDateTime.slice(0, 10)
    let time = lastMessageDateTime.slice(11, 16);
    return (

        <div onClick={() => { props.setOtherUserId(props.otherUserId) }}>
            <li className="left clearfix">
                <span className="chat-img pull-left">
                    <img
                        src={require("../utils/user.png")}
                        alt="User Avatar"
                        className="img-circle"
                    />
                </span>
                <div className="chat-body clearfix">
                    <div className="header_sec">
                        <strong className="primary-font">{props.chatroom.other_user.nickname}</strong>{" "}

                    </div>
                    <div className="contact_sec">
                        <strong className="pull-right">{time}, {date}</strong>

                        {/* <span className="badge pull-right">3</span> */}
                    </div>
                    <div className="contact_sec">
                        <strong className="primary-font">{lastMessage}</strong>
                        {/* <span className="badge pull-right">3</span> */}
                    </div>
                </div>
            </li>
        </div>
    )
}

export default PersonBox