# /// script
# requires-python = ">=3.10"
# dependencies = ["iterm2"]
# ///

import asyncio
import threading
from http.server import ThreadingHTTPServer

import iterm2

try:
    from .bridge_handler import BridgeHandler
except ImportError:
    from bridge_handler import BridgeHandler

HOST = "127.0.0.1"
PORT = 6124


class BridgeServer(ThreadingHTTPServer):
    def __init__(self, connection, loop):
        super().__init__((HOST, PORT), BridgeHandler)
        self.connection = connection
        self.loop = loop


async def main(connection):
    loop = asyncio.get_event_loop()
    app = await iterm2.async_get_app(connection)
    connection.app = app
    server = BridgeServer(connection, loop)
    thread = threading.Thread(target=server.serve_forever, daemon=True)
    thread.start()
    print("iTermBridge: listening on {}:{}".format(HOST, PORT))

    try:
        while True:
            await asyncio.sleep(1)
    except asyncio.CancelledError:
        pass
    finally:
        server.shutdown()
        print("iTermBridge: stopped")


try:
    iterm2.run_until_complete(main)
except KeyboardInterrupt:
    pass
