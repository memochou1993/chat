import React from "react";
import "./style.scss";

const History = (props) => {
  return (
    <div className="History">
      {
        props.messages.map((message, index) => (
          <p
            key={ index }
          >
            { message.data }
          </p>
        ))
      }
    </div>
  );
};

export default History;
