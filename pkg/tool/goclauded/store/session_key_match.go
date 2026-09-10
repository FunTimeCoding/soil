package store

func sessionKeyMatch(
	sessionIdentifier string,
	callsign string,
) (string, []any) {
	if callsign == "" {
		return "session_identifier = ?", []any{sessionIdentifier}
	}

	return `(session_identifier = ?
		OR ((session_identifier IS NULL OR session_identifier = '')
		AND callsign = ?))`, []any{sessionIdentifier, callsign}
}
