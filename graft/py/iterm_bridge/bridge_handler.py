import asyncio
import json
import re
from http.server import BaseHTTPRequestHandler

try:
    from . import bridge_core
except ImportError:
    import bridge_core

MALFORMED = 400
BROKEN = 500


class BridgeHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        try:
            self._get()
        except Exception as e:
            self._error_response(BROKEN, str(e))

    def do_POST(self):
        try:
            self._post()
        except Exception as e:
            self._error_response(BROKEN, str(e))

    def _get(self):
        path, _, query_string = self.path.partition("?")

        if path == "/sessions":
            result = self._run(bridge_core.list_sessions(self.server.connection))
            self._json_response({"sessions": result})
            return

        match = re.match(r"^/sessions/([^/]+)/screen$", path)

        if match:
            session_id = match.group(1)
            result, failure = self._run(
                bridge_core.read_screen(self.server.connection, session_id)
            )
            self._answer(result, failure)
            return

        match = re.match(r"^/sessions/([^/]+)/history$", path)

        if match:
            session_id = match.group(1)
            count = 50

            if query_string:
                for pair in query_string.split("&"):
                    key, _, value = pair.partition("=")

                    if key == "count":
                        count = int(value)

            result, failure = self._run(
                bridge_core.read_history(
                    self.server.connection, session_id, count
                )
            )
            self._answer(result, failure)
            return

        self._error_response(bridge_core.NOT_FOUND, "not found")

    def _post(self):
        match = re.match(r"^/sessions/([^/]+)/send$", self.path)

        if match:
            session_id = match.group(1)
            body = self._read_body()

            if body is None:
                return

            text = body.get("text", "")

            if not text:
                self._error_response(bridge_core.UNUSABLE, "text is required")
                return

            result, failure = self._run(
                bridge_core.send_text(self.server.connection, session_id, text)
            )
            self._answer(result, failure)
            return

        match = re.match(r"^/sessions/([^/]+)/key$", self.path)

        if match:
            session_id = match.group(1)
            body = self._read_body()

            if body is None:
                return

            key = body.get("key", "")

            if not key:
                self._error_response(bridge_core.UNUSABLE, "key is required")
                return

            result, failure = self._run(
                bridge_core.send_key(self.server.connection, session_id, key)
            )
            self._answer(result, failure)
            return

        match = re.match(r"^/tabs/([^/]+)/title$", self.path)

        if match:
            tab_id = match.group(1)
            body = self._read_body()

            if body is None:
                return

            title = body.get("title", "")

            if not title:
                self._error_response(bridge_core.UNUSABLE, "title is required")
                return

            result, failure = self._run(
                bridge_core.set_tab_title(self.server.connection, tab_id, title)
            )
            self._answer(result, failure)
            return

        match = re.match(r"^/sessions/([^/]+)/color$", self.path)

        if match:
            session_id = match.group(1)
            body = self._read_body()

            if body is None:
                return

            red = body.get("red", 0)
            green = body.get("green", 0)
            blue = body.get("blue", 0)
            result, failure = self._run(
                bridge_core.set_tab_color(
                    self.server.connection, session_id, red, green, blue
                )
            )
            self._answer(result, failure)
            return

        if self.path == "/tabs":
            result, failure = self._run(
                bridge_core.create_tab(self.server.connection)
            )
            self._answer(result, failure, 201)
            return

        self._error_response(bridge_core.NOT_FOUND, "not found")

    def _answer(self, result, failure, status=200):
        if failure:
            self._error_response(failure.status, failure.message)
            return

        self._json_response(result, status)

    def _run(self, coroutine):
        return asyncio.run_coroutine_threadsafe(
            coroutine, self.server.loop
        ).result()

    def _read_body(self):
        length = int(self.headers.get("Content-Length", 0))

        if length == 0:
            self._error_response(MALFORMED, "empty request body")
            return None

        try:
            return json.loads(self.rfile.read(length))
        except json.JSONDecodeError:
            self._error_response(MALFORMED, "invalid JSON")
            return None

    def _json_response(self, data, status=200):
        body = json.dumps(data, ensure_ascii=False).encode("utf-8")
        self.send_response(status)
        self.send_header("Content-Type", "application/json")
        self.send_header("Content-Length", str(len(body)))
        self.end_headers()
        self.wfile.write(body)

    def _error_response(self, status, message):
        self._json_response({"error": message}, status)

    def log_message(self, format, *arguments):
        pass
