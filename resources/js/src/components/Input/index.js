import React from 'react';
import './style.scss';

const Input = (props) => {
  return (
    <div className="Input">
      <form
        onSubmit={props.submit}
      >
        <input
          value={props.value}
          onChange={props.onChange}
        />
      </form>
    </div>
  );
};

export default Input;
