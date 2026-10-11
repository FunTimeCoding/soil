import threading
from http.server import ThreadingHTTPServer

try:
    from .bridge_handler import BridgeHandler
except ImportError:
    from bridge_handler import BridgeHandler

HOST = "127.0.0.1"
PORT = 6123


class BridgeServer:
    def __init__(self):
        self._server = None
        self._thread = None

    def start(self):
        self._server = ThreadingHTTPServer((HOST, PORT), BridgeHandler)
        self._thread = threading.Thread(
            target=self._server.serve_forever, daemon=True
        )
        self._thread.start()
        print("SublimeBridge: listening on {}:{}".format(HOST, PORT))

    def stop(self):
        if self._server:
            self._server.shutdown()
            self._server = None
            self._thread = None
            print("SublimeBridge: stopped")
