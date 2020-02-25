import React, { useEffect, useState } from 'react';
import Header from './components/Header';
import History from './components/History';
import Input from './components/Input';
import { connect, send } from './api';

const App = () => {
  const [messages, setMessages] = useState([]);

  useEffect(() => {
    connect((message) => {
      setMessages(prevState => {
        return [...prevState, message];
      });
    });
  }, []);

  const handleClick = () => {
    send('Hello!');
  };

  const handleKeyDown = (event) => {
    if (event.keyCode === 13) {
      send(event.target.value);
      event.target.value = "";
    }
  };

  return (
    <div className="App">
      <Header />
      <button
        onClick={handleClick}
      >
        Send
      </button>
      <History
        messages={messages}
      />
      <Input
        onKeyDown={handleKeyDown}
      />
    </div>
  );
};

export default App;
