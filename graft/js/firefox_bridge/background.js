const WEBSOCKET_URL = "ws://127.0.0.1:6125";
const RECONNECT_INTERVAL = 3000;

let socket = null;
let connecting = false;

function connect() {
  if (connecting) {
    return;
  }

  connecting = true;

  try {
    socket = new WebSocket(WEBSOCKET_URL);
  } catch (error) {
    connecting = false;
    scheduleReconnect();
    return;
  }

  socket.onopen = function () {
    connecting = false;
    console.log("FirefoxBridge: connected to", WEBSOCKET_URL);
  };

  socket.onclose = function () {
    console.log("FirefoxBridge: disconnected");
    socket = null;
    connecting = false;
    scheduleReconnect();
  };

  socket.onerror = function () {
    // onclose will fire after onerror, so don't reconnect here
  };

  socket.onmessage = async function (event) {
    let request;

    try {
      request = JSON.parse(event.data);
    } catch (error) {
      return;
    }

    try {
      const result = await dispatch(request.method, request.params);
      socket.send(JSON.stringify({ result: result, id: request.id }));
    } catch (error) {
      socket.send(
        JSON.stringify({ error: error.message, id: request.id })
      );
    }
  };
}

function scheduleReconnect() {
  setTimeout(connect, RECONNECT_INTERVAL);
}

connect();
