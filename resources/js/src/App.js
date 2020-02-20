import React, { useEffect, useState } from 'react';
import Header from './components/Header';
import History from './components/History';
import { connect, sendMessage } from './api';

const App = () => {
  const [messages, setMessages] = useState([]);

  useEffect(() => {
    connect((message) => {
      setMessages(prevState => {
        return [...prevState, message];
      });
    });
  }, []);

  const send = () => {
    sendMessage('Hello!');
  };

  return (
    <div className="App">
      <Header />
      <button
        onClick={ send }
      >
        Send
      </button>
      <History
        messages={ messages }
      />
    </div>
  );
};

export default App;
