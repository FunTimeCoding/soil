import iterm2


async def list_sessions(connection):
    app = await iterm2.async_get_app(connection)
    result = []

    for window in app.windows:
        window_id = window.window_id
        window_title = await window.async_get_variable("title") or ""

        for tab in window.tabs:
            tab_id = tab.tab_id
            tab_title = await tab.async_get_variable("title") or ""

            for session in tab.sessions:
                result.append(await _session_summary(
                    session, window_id, window_title, tab_id, tab_title,
                ))

    return result


async def read_screen(connection, session_id):
    session = _find_session(connection, session_id)

    if session is None:
        return None, "session not found"

    contents = await session.async_get_screen_contents()
    lines = []

    for i in range(contents.number_of_lines):
        lines.append(contents.line(i).string)

    cursor = contents.cursor_coord

    return {
        "session_id": session_id,
        "lines": lines,
        "cursor_x": cursor.x,
        "cursor_y": cursor.y,
        "columns": contents.number_of_lines_above_screen + contents.number_of_lines,
    }, None


async def read_history(connection, session_id, count):
    session = _find_session(connection, session_id)

    if session is None:
        return None, "session not found"

    contents = await session.async_get_screen_contents()
    total_above = contents.number_of_lines_above_screen
    screen_lines = contents.number_of_lines

    if count <= 0:
        count = 50

    start = max(0, total_above + screen_lines - count)
    length = min(count, total_above + screen_lines - start)
    line_info = await session.async_get_contents(start, length)
    lines = []

    for line in line_info:
        lines.append(line.string)

    return {
        "session_id": session_id,
        "lines": lines,
        "start_line": start,
        "total_scrollback": total_above,
    }, None


async def send_text(connection, session_id, text):
    session = _find_session(connection, session_id)

    if session is None:
        return None, "session not found"

    await session.async_send_text(text)

    return {"session_id": session_id, "sent": text}, None


KEY_MAP = {
    "enter": "\r",
    "tab": "\t",
    "escape": "\x1b",
    "ctrl+c": "\x03",
    "ctrl+d": "\x04",
    "ctrl+z": "\x1a",
    "ctrl+l": "\x0c",
    "ctrl+a": "\x01",
    "ctrl+e": "\x05",
    "ctrl+r": "\x12",
    "ctrl+w": "\x17",
    "ctrl+u": "\x15",
    "up": "\x1b[A",
    "down": "\x1b[B",
    "right": "\x1b[C",
    "left": "\x1b[D",
    "backspace": "\x7f",
    "delete": "\x1b[3~",
}


async def send_key(connection, session_id, key):
    session = _find_session(connection, session_id)

    if session is None:
        return None, "session not found"

    sequence = KEY_MAP.get(key.lower())

    if sequence is None:
        return None, "unknown key: {}. valid keys: {}".format(
            key, ", ".join(sorted(KEY_MAP.keys()))
        )

    await session.async_send_text(sequence)

    return {"session_id": session_id, "key": key}, None


def _find_tab(connection, tab_id):
    app = connection.app

    if app is None:
        return None

    for window in app.windows:
        for tab in window.tabs:
            if tab.tab_id == tab_id:
                return tab

    return None


async def set_tab_title(connection, tab_id, title):
    tab = _find_tab(connection, tab_id)

    if tab is None:
        return None, "tab not found"

    await tab.async_set_title(title)

    return {"tab_id": tab_id, "title": title}, None


async def set_tab_color(connection, session_id, red, green, blue):
    session = _find_session(connection, session_id)

    if session is None:
        return None, "session not found"

    change = iterm2.LocalWriteOnlyProfile()
    change.set_tab_color(iterm2.Color(red, green, blue))
    change.set_use_tab_color(True)
    await session.async_set_profile_properties(change)

    return {"session_id": session_id, "red": red, "green": green, "blue": blue}, None


async def create_tab(connection):
    app = await iterm2.async_get_app(connection)
    window = app.current_window

    if window is None:
        return None, "no window"

    tab = await window.async_create_tab()
    session = tab.current_session
    tab_id = tab.tab_id
    tab_title = await tab.async_get_variable("title") or ""
    window_id = window.window_id
    window_title = await window.async_get_variable("title") or ""

    return await _session_summary(
        session, window_id, window_title, tab_id, tab_title,
    ), None


def _find_session(connection, session_id):
    app = connection.app

    if app is None:
        return None

    for window in app.windows:
        for tab in window.tabs:
            for session in tab.sessions:
                if session.session_id == session_id:
                    return session

    return None


async def _session_summary(session, window_id, window_title, tab_id, tab_title):
    job_name = await session.async_get_variable("jobName") or ""
    command_line = await session.async_get_variable("commandLine") or ""
    tty = await session.async_get_variable("tty") or ""
    path = await session.async_get_variable("path") or ""
    columns = await session.async_get_variable("columns") or 0
    rows = await session.async_get_variable("rows") or 0

    return {
        "session_id": session.session_id,
        "window_id": window_id,
        "window_title": window_title,
        "tab_id": tab_id,
        "tab_title": tab_title,
        "job_name": job_name,
        "command_line": command_line,
        "path": path,
        "tty": tty,
        "columns": columns,
        "rows": rows,
    }
