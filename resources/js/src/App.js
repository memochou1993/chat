import React, { useEffect } from 'react';
import './App.css';
import { connect, sendMessage } from "./api";

const App = () => {
  useEffect(() => {
    connect();
  }, []);
  const send = () => {
    sendMessage('Hello!');
  };
  return (
    <div className="App">
      <button onClick={send}>Hit</button>
    </div>
  );
};

export default App;
