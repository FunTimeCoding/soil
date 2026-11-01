package pulse_call

func New(
	session string,
	body string,
) *Call {
	return &Call{Session: session, Body: body}
}
