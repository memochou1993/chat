import React from 'react';
import './style.scss';

const History = (props) => {
  return (
    <div
      className="History"
    >
      {
        props.messages.map((message, index) => {
          return (
            <div
              key={index}
              className="alert alert-primary"
            >
              {message.data}
            </div>
          );
        })
      }
    </div>
  );
};

export default History;
