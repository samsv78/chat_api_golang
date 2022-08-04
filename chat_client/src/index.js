import React from 'react';
import ReactDOM from 'react-dom/client';

import App from './App';
window.IPADRESS = '192.168.1.5'
const root = ReactDOM.createRoot(document.getElementById('root'));
window.addEventListener("beforeunload", (ev) => {
  ev.preventDefault();
  return ev.returnValue = 'Are you sure you want to close?';
});
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

