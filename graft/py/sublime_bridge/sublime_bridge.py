import sublime
import sublime_plugin

try:
    from .bridge_server import BridgeServer
except ImportError:
    from bridge_server import BridgeServer

_server = None


def plugin_loaded():
    global _server
    _server = BridgeServer()
    _server.start()


def plugin_unloaded():
    global _server

    if _server:
        _server.stop()
        _server = None


class SublimeBridgeReplaceCommand(sublime_plugin.TextCommand):
    def run(self, edit, regions, new_string):
        for region in regions:
            self.view.replace(edit, sublime.Region(region[0], region[1]), new_string)
