import React from 'react';
import './style.scss';

const History = (props) => {
  return (
    <div
      className="History"
    >
      {
        props.messages.map((message, index) => {
          if (!message.roomId) {
            return (
              <div
                key={index}
                className="row"
              >
                <div
                  className="col-md-6 offset-md-3"
                >
                  <div
                    className="alert alert-info text-center my-3 py-1"
                  >
                    {message.body}
                  </div>
                </div>
              </div>
            );
          }

          if (message.clientId === props.self) {
            return (
              <div
                key={index}
                className="row"
              >
                <div
                  className="col-md-6 offset-md-6"
                >
                  <div
                    className="alert alert-primary my-3"
                  >
                    {message.body}
                  </div>
                </div>
              </div>
            );
          }

          return (
            <div
              key={index}
              className="row"
            >
              <div
                className="col-md-6"
              >
                <div
                  className="alert alert-primary my-3"
                >
                  {message.body}
                </div>
              </div>
            </div>
          );
        })
      }
    </div>
  );
};

export default History;
