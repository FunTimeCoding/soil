package queue

func NewImmediate(
	sessionIdentifier string,
	callsign string,
	kind string,
	body string,
) *Entry {
	result := New(sessionIdentifier, callsign, kind, body)
	result.Immediate = true

	return result
}
