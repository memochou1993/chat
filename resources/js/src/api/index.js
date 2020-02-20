const socket = new WebSocket('ws://localhost:8082');

const connect = () => {
  console.log("Attempting Connection...");

  socket.onopen = (event) => {
    console.log("Connection Opened: ", event);
  };

  socket.onmessage = (message) => {
    console.log(message);
  };

  socket.onclose = (event) => {
    console.log("Connection Closed: ", event);
  };

  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
  };
};

const sendMessage = (message) => {
  console.log("Sending Message: ", message);
  socket.send(message);
};

export {
    connect,
    sendMessage,
};
