package top_entry

func New(
	environment string,
	target string,
	state string,
	live bool,
) *Entry {
	return &Entry{
		Environment: environment,
		Target:      target,
		State:       state,
		Live:        live,
	}
}
