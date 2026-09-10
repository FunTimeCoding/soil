package notification

func New(
	sessionIdentifier string,
	callsign string,
	source string,
	body string,
) *Notification {
	return &Notification{
		SessionIdentifier: sessionIdentifier,
		Callsign:          callsign,
		Source:            source,
		Body:              body,
	}
}
