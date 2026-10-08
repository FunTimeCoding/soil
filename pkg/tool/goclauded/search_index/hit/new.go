package hit

func New(
	identifier string,
	role string,
	kind string,
	at string,
	snippet string,
) *Hit {
	return &Hit{
		Identifier: identifier,
		Role:       role,
		Kind:       kind,
		At:         at,
		Snippet:    snippet,
	}
}
