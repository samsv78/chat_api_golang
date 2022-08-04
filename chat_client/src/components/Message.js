import React from 'react'

const Message = (props) => {
    // console.log(props.userId)
    // let a = props.userId === props.message.receiver.id
    // console.log(a)
    let datetime = props.message.send_date
    let date = datetime.slice(0, 10)
    let time = datetime.slice(11, 16);
    return (
        <>
            {
                props.userId === props.message.sender.id
                    ?
                    <li className="left clearfix">
                        <span className="chat-img1 pull-left">
                            <img
                                src={require("../utils/user.png")}
                                alt="User Avatar"
                                className="img-circle"
                            />
                        </span>
                        <div className="chat-body1 clearfix">
                            <p>
                                {props.message.text}
                            </p>
                            <div className="chat_time pull-right">{time}, {date}</div>
                        </div>
                    </li>
                    :
                    <li className="left clearfix admin_chat">
                        <span className="chat-img1 pull-right">
                            <img
                                src={require("../utils/user.png")}
                                alt="User Avatar"
                                className="img-circle"
                            />
                        </span>
                        <div className="chat-body1 clearfix">
                            <p>
                                {props.message.text}
                            </p>
                            <div className="chat_time pull-left">{time}, {date}</div>
                        </div>
                    </li>
            }
        </>
    )
}

export default Message