import sublime


def list_views():
    result = []

    for window in sublime.windows():
        for view in window.views():
            result.append(_view_summary(view, window))

    return result


def read_view(view_id):
    view, window = _find_view(view_id)

    if view is None:
        return None

    return _view_detail(view, window)


def edit_view(view_id, old_string, new_string, replace_all):
    view, window = _find_view(view_id)

    if view is None:
        return None, "view not found"

    if replace_all:
        regions = view.find_all(old_string, sublime.LITERAL)
    else:
        regions = [view.find(old_string, 0, sublime.LITERAL)]

    if not regions or regions[0].empty():
        return None, "old_string not found in buffer"

    if not replace_all and len(view.find_all(old_string, sublime.LITERAL)) > 1:
        return None, "old_string matches multiple times; use replace_all"

    view.run_command(
        "sublime_bridge_replace",
        {
            "regions": [[r.a, r.b] for r in reversed(regions)],
            "new_string": new_string,
        },
    )

    return _view_detail(view, window), None


def create_view(title, content, syntax):
    window = sublime.active_window()
    view = window.new_file()
    view.set_scratch(True)

    if title:
        view.set_name(title)

    if syntax and hasattr(sublime, 'find_syntax'):
        view.assign_syntax(sublime.find_syntax(syntax))

    if content:
        view.run_command(
            "sublime_bridge_replace",
            {"regions": [[0, 0]], "new_string": content},
        )

    return _view_detail(view, window)


def save_view(view_id, file_path):
    view, window = _find_view(view_id)

    if view is None:
        return None, "view not found"

    if file_path:
        view.retarget(file_path)
        view.set_scratch(False)

    view.run_command("save")

    return _view_detail(view, window), None


def close_view(view_id):
    view, window = _find_view(view_id)

    if view is None:
        return None, "view not found"

    result = _view_summary(view, window)
    view.set_scratch(True)
    window.focus_view(view)
    window.run_command("close_file")

    return result, None


def open_file(file_path):
    window = sublime.active_window()
    view = window.open_file(file_path)

    return _view_detail(view, window)


def _find_view(view_id):
    for window in sublime.windows():
        for view in window.views():
            if view.buffer_id() == view_id:
                return view, window

    return None, None


def _view_summary(view, window):
    text = view.substr(sublime.Region(0, min(view.size(), 200)))

    return {
        "view_id": view.buffer_id(),
        "window_id": window.id(),
        "file_path": view.file_name(),
        "title": _view_title(view),
        "is_dirty": view.is_dirty(),
        "syntax": _syntax_name(view),
        "preview": text,
    }


def _view_detail(view, window):
    text = view.substr(sublime.Region(0, view.size()))

    return {
        "view_id": view.buffer_id(),
        "window_id": window.id(),
        "file_path": view.file_name(),
        "title": _view_title(view),
        "is_dirty": view.is_dirty(),
        "syntax": _syntax_name(view),
        "text": text,
    }


def _view_title(view):
    if view.file_name():
        return view.file_name().split("/")[-1]

    return view.name() or "untitled"


def _syntax_name(view):
    syntax = view.syntax()

    if syntax:
        return syntax.name

    return ""
