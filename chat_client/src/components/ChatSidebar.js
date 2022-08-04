import React, { useState, useEffect } from 'react';
import PersonBox from './PersonBox'
import SearchBox from './SearchBox'


const ChatSidebar = (props) => {
    const [chatrooms, setChatrooms] = useState([]);
    useEffect(function effectFunction() {

        async function getChatrooms() {
            const response = await fetch('http://' + window.IPADRESS + ':8080/message/chatrooms', {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer ' + props.token,
                }
            })
            let json = await response.json()
            setChatrooms(json)
        }
        getChatrooms();
    }, [props.token, props.sending, props.receiving]);
    // console.log(chatrooms)
    return (
        <>
            <div className='col-sm-2 chat_sidebar'>
                <div className='row'>
                    <SearchBox />
                    <div className='member_list'>
                        <ul className='list-unstyled'>
                            {
                                chatrooms.map((chatroom) => (
                                    <PersonBox key={chatroom.user.id.toString() + chatroom.other_user.id.toString()} chatroom={chatroom} setOtherUserId={props.setOtherUserId} otherUserId={chatroom.other_user.id} />
                                ))
                            }
                            {/* {
                                tasks.map((task, index) => (
                                    <Task key={index} task={task} onDelete={onDelete} onToggle={onToggle} />
                                ))
                            } */}

                        </ul>
                    </div>
                </div>
            </div>
        </>
    )
}

export default ChatSidebar