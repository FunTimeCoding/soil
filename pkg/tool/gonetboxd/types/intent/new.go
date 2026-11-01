package intent

func New(
	kind string,
	name string,
	status string,
) *Intent {
	return &Intent{Kind: kind, Name: name, Status: status}
}
