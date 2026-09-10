package queue

func New(
	sessionIdentifier string,
	callsign string,
	kind string,
	body string,
) *Entry {
	return &Entry{
		SessionIdentifier: sessionIdentifier,
		Callsign:          callsign,
		Kind:              kind,
		Body:              body,
	}
}
