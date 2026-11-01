package notify_call

func New(
	callsign string,
	source string,
	body string,
) *Call {
	return &Call{Callsign: callsign, Source: source, Body: body}
}
