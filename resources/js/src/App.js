import React, { useEffect, useState } from 'react';
import Header from './components/Header';
import History from './components/History';
import Input from './components/Input';
import { connect, send } from './api';
import 'bootstrap/dist/css/bootstrap.min.css';
import './style.scss';

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
      <div
        className="container-fulid"
      >
        <div
          id="header"
        >
          <Header />
        </div>
        <div
          id="body"
        >
          <History
            messages={messages}
          />
        </div>
        <div
          id="footer"
        >
          <Input
            submit={handleSubmit}
            value={message}
            onChange={handleChange}
          />
        </div>
      </div>
    </div>
  );
};

export default App;
