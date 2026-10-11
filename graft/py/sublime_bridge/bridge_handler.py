import json
import re
from http.server import BaseHTTPRequestHandler

try:
    from . import bridge_core
except ImportError:
    import bridge_core


class BridgeHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/views":
            self._json_response({"views": bridge_core.list_views()})
            return

        match = re.match(r"^/views/(\d+)$", self.path)

        if match:
            view_id = int(match.group(1))
            result = bridge_core.read_view(view_id)

            if result is None:
                self._error_response(404, "view not found")
                return

            self._json_response(result)
            return

        self._error_response(404, "not found")

    def do_POST(self):
        if self.path == "/views":
            body = self._read_body_optional()
            title = body.get("title", "") if body else ""
            content = body.get("content", "") if body else ""
            syntax = body.get("syntax", "") if body else ""
            result = bridge_core.create_view(title, content, syntax)
            self._json_response(result, 201)
            return

        match = re.match(r"^/views/(\d+)/save$", self.path)

        if match:
            view_id = int(match.group(1))
            body = self._read_body_optional()
            file_path = body.get("file_path", "") if body else ""
            result, error = bridge_core.save_view(view_id, file_path)

            if error == "view not found":
                self._error_response(404, error)
                return

            if error:
                self._error_response(400, error)
                return

            self._json_response(result)
            return

        if self.path == "/open":
            body = self._read_body()

            if body is None:
                return

            file_path = body.get("file_path", "")

            if not file_path:
                self._error_response(400, "file_path is required")
                return

            result = bridge_core.open_file(file_path)
            self._json_response(result)
            return

        self._error_response(404, "not found")

    def do_PUT(self):
        match = re.match(r"^/views/(\d+)$", self.path)

        if not match:
            self._error_response(404, "not found")
            return

        view_id = int(match.group(1))
        body = self._read_body()

        if body is None:
            return

        old_string = body.get("old_string", "")
        new_string = body.get("new_string", "")
        replace_all = body.get("replace_all", False)

        if not old_string:
            self._error_response(400, "old_string is required")
            return

        result, error = bridge_core.edit_view(
            view_id, old_string, new_string, replace_all
        )

        if error == "view not found":
            self._error_response(404, error)
            return

        if error:
            self._error_response(400, error)
            return

        self._json_response(result)

    def do_DELETE(self):
        match = re.match(r"^/views/(\d+)$", self.path)

        if not match:
            self._error_response(404, "not found")
            return

        view_id = int(match.group(1))
        result, error = bridge_core.close_view(view_id)

        if error == "view not found":
            self._error_response(404, error)
            return

        if error:
            self._error_response(400, error)
            return

        self._json_response(result)

    def _read_body(self):
        length = int(self.headers.get("Content-Length", 0))

        if length == 0:
            self._error_response(400, "empty request body")
            return None

        try:
            return json.loads(self.rfile.read(length))
        except json.JSONDecodeError:
            self._error_response(400, "invalid JSON")
            return None

    def _read_body_optional(self):
        length = int(self.headers.get("Content-Length", 0))

        if length == 0:
            return {}

        try:
            return json.loads(self.rfile.read(length))
        except json.JSONDecodeError:
            return {}

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
