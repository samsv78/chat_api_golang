import Message from './Message'
import React, { useState, useEffect, useRef } from 'react';
const MessageSection = (props) => {
    const bottomRef = useRef(null);
    const [chatroom, setChatroom] = useState([]);
    useEffect(function effectFunction() {

        async function getChatroom() {
            if (props.otherUserId !== 0) {
                const response = await fetch('http://' + window.IPADRESS + ':8080/message/chatrooms/' + props.otherUserId, {
                    method: 'GET',
                    headers: {
                        'Authorization': 'Bearer ' + props.token,
                    }
                })
                let json = await response.json()
                setChatroom(json)
            }

        }
        getChatroom();
        bottomRef.current?.scrollIntoView({ behavior: 'smooth' });
    }, [props.token, props.otherUserId, props.sending, props.receiving]);
    let messages = []
    if (chatroom.hasOwnProperty('messages')) {
        messages = chatroom.messages
    }
    // console.log(messages)
    return (
        <>
            <div className="col-sm-10 message_section">
                <div className="row">
                    <div className="new_message_head">
                        <div className="pull-left">
                            <button>
                                <i className="fa fa-plus-square-o" aria-hidden="true" /> NewMessage
                            </button>
                        </div>
                        <div className="pull-right btn btn-danger">
                            <button onClick={props.logout} className='btn'> logout </button>
                        </div>
                    </div>
                    {/*new_message_head*/}
                    <div className="chat_area">
                        <ul className="list-unstyled">
                            {
                                messages.length > 0 ?
                                    messages.map((message) => (
                                        < Message key={message.message_id} message={message} userId={chatroom.user.id} />
                                    ))
                                    :
                                    <></>

                            }
                        </ul>
                    </div>
                    {/*chat_area*/}
                    <div className="message_write">
                        <textarea
                            className="form-control"
                            placeholder="type a message"
                            defaultValue={""}
                            id="messageText"
                        />
                        <div className="clearfix" />
                        <div className="chat_bottom">
                            <button onClick={() => props.send(chatroom.other_user.id)} className="pull-right btn btn-success">
                                Send
                            </button>
                        </div>

                    </div>

                </div>

            </div>
        </>
    )
}

export default MessageSection