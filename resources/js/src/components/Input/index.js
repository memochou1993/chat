import React from "react";
import "./style.scss";

const Input = (props) => {
  return (
    <div className="Input">
      <input
        onKeyDown={props.onKeyDown}
      />
    </div>
  );
};

export default Input;
