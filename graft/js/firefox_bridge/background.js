const HOST_NAME = "org.funtimecoding.firefox_relay";
const RECONNECT_INTERVAL = 3000;

let port = null;

console.log(
  "FirefoxBridge: background loaded, permissions",
  browser.runtime.getManifest().permissions.join(",")
);

function connect() {
  if (port) {
    return;
  }

  try {
    port = browser.runtime.connectNative(HOST_NAME);
  } catch (error) {
    console.error("FirefoxBridge: connectNative threw:", error.message);
    scheduleReconnect();

    return;
  }

  console.log("FirefoxBridge: port opened to", HOST_NAME);

  port.onDisconnect.addListener(function (p) {
    const fail = p.error ? p.error.message : "";
    console.error("FirefoxBridge: port disconnected:", fail || "no reason");
    port = null;
    scheduleReconnect();
  });

  port.onMessage.addListener(async function (request) {
    let reply;

    try {
      reply = { result: await dispatch(request.method, request.params) };
    } catch (error) {
      reply = { error: error.message };
    }

    if (!port) {
      return;
    }

    reply.id = request.id;
    port.postMessage(reply);
  });
}

function scheduleReconnect() {
  setTimeout(connect, RECONNECT_INTERVAL);
}

connect();
