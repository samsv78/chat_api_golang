import React, { useState } from 'react';
import './App.css';
import ChatSidebar from './components/ChatSidebar';
import Login from './components/Login';
import MessageSection from './components/MessageSection';


function App() {
  const [socket, setSocket] = useState(null);
  const [isLoggedIn, setIsLoggenIn] = useState(false);
  const [token, setToken] = useState(null);
  const [otherUserId, setOtherUserId] = useState(0);
  const [sending, setSending] = useState(false);
  const [receiving, setReceinving] = useState(false);

  const login = async () => {
    let email = document.getElementById('email').value;
    let password = document.getElementById('password').value;
    let loginRequest = {
      email: email,
      password: password
    }
    const response = await fetch('http://' + window.IPADRESS + ':8080/login', {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
      },
      body: JSON.stringify(loginRequest),
    })
    const token = await response.json()
    initWebSoket(token)
  }

  const initWebSoket = (token) => {
    let socket = new WebSocket('ws://' + window.IPADRESS + ':8080/ws')
    console.log('attempting websocket connection')
    socket.onopen = () => {
      console.log('websocket connected!')
      socket.send(token)
    }

    setSocket(socket)
    setIsLoggenIn(true)
    setToken(token)
  }
  const logout = () => {
    socket.close()
    setSocket(null)
    setIsLoggenIn(false)
    setToken(null)
    setOtherUserId(0)

  }

  const send = async (receiverId) => {
    let text = document.getElementById('messageText').value;
    if (text === '') {
      return
    }
    document.getElementById('messageText').value = '';
    let sendMessageRequest = {
      text: text,
      receiver_id: receiverId
    }
    const response = await fetch('http://' + window.IPADRESS + ':8080/message', {
      method: 'POST',
      headers: {
        'Authorization': 'Bearer ' + token,
      },
      body: JSON.stringify(sendMessageRequest),
    })
    const data = await response.json()
    console.log(data)
    if (sending) {
      setSending(false)
    } else {
      setSending(true)
    }

  }

  if (!isLoggedIn) {
    return (
      <Login login={login} />
      // <>
      //   <div>The user is <b>{isLoggedIn ? 'currently' : 'not'}</b> logged in.</div>
      //   <br />
      //   <br />
      //   <input type='text' id='email' placeholder='email' />
      //   <br />
      //   <input type='text' id='password' placeholder='password' />
      //   <br />
      //   <br />
      //   <button onClick={login} style={{ backgroundColor: 'green' }} className='btn'> login </button>
      //   <br />
      //   <br />
      // </>
    )
  } else {
    socket.addEventListener('message', function (event) {
      console.log('Message from server ', event.data);
      if (event.data === 'RELOAD') {
        if (receiving) {
          setReceinving(false)
        } else {
          setReceinving(true)
        }
      }
    });
    return (
      <>

        <div className='main_section'>
          <div className='container'>
            <div className='chat_container'>

              <ChatSidebar token={token} setOtherUserId={setOtherUserId} sending={sending} receiving={receiving} />

              <MessageSection token={token} otherUserId={otherUserId} logout={logout} send={send} sending={sending} receiving={receiving} />

            </div>
          </div>
        </div>
      </>
    )
  }

}

export default App