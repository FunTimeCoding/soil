package queue

func New(
	callsign string,
	kind string,
	body string,
) *Entry {
	return &Entry{Callsign: callsign, Kind: kind, Body: body}
}
