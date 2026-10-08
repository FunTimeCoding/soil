package entry

func New(
	identifier string,
	session string,
	turn string,
	role string,
	kind string,
	at string,
	body string,
) *Entry {
	return &Entry{
		Identifier: identifier,
		Session:    session,
		Turn:       turn,
		Role:       role,
		Kind:       kind,
		At:         at,
		Body:       body,
	}
}
