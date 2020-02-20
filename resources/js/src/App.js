import React, { useEffect } from 'react';
import Header from './components/Header';
import { connect, sendMessage } from './api';

const App = () => {
  useEffect(() => {
    connect();
  }, []);

  const send = () => {
    sendMessage('Hello!');
  };

  return (
    <div className="App">
      <Header />
      <button onClick={send}>Hit</button>
    </div>
  );
};

export default App;
