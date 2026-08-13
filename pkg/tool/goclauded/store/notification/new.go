package notification

func New(
	callsign string,
	source string,
	body string,
) *Notification {
	return &Notification{Callsign: callsign, Source: source, Body: body}
}
