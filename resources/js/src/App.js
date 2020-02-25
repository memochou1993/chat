import React, { useEffect, useState } from 'react';
import Header from './components/Header';
import History from './components/History';
import Input from './components/Input';
import { connect, send } from './api';

const App = () => {
  const [message, setMessage] = useState('');
  const [messages, setMessages] = useState([]);

  useEffect(() => {
    connect((state) => {
      setMessages((prevState) => {
        return [...prevState, state];
      });
    });
  }, []);

  const handleSubmit = (event) => {
    send(message);
    setMessage('');
    event.preventDefault();
  };

  const handleChange = (event) => {
    setMessage(event.target.value);
  };

  return (
    <div className="App">
      <Header />
      <History
        messages={messages}
      />
      <Input
        submit={handleSubmit}
        value={message}
        onChange={handleChange}
      />
    </div>
  );
};

export default App;
