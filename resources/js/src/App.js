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
  const [self, setSelf] = useState('');

  useEffect(() => {
    connect((state) => {
      setMessages((prevState) => {
        const data = JSON.parse(state.data);

        if (!data.roomId) {
          setSelf(data.clientId);
        }

        return [...prevState, data];
      });
    });
  }, []);

  const handleSubmit = (event) => {
    if (message) {
      send(message);
      setMessage('');
    }
    event.preventDefault();
  };

  const handleChange = (event) => {
    setMessage(event.target.value);
  };

  return (
    <div className="App">
      <div
        className="container-fulid bg-light"
      >
        <div
          id="header"
        >
          <Header />
        </div>
        <div
          id="body"
          className="overflow-auto"
        >
          <History
            self={self}
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
