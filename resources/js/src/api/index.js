const socket = new WebSocket('ws://localhost:8080');

const connect = () => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected!");
  };

  socket.onmessage = message => {
    console.log(message);
  };

  socket.onclose = event => {
    console.log("Connection Closed: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

const sendMessage = message => {
  console.log("Sending Message: ", message);
  socket.send(message);
};

export {
    connect,
    sendMessage,
};
