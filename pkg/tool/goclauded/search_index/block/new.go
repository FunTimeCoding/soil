package block

func New(
	identifier string,
	role string,
	kind string,
	at string,
	text string,
) *Block {
	return &Block{
		Identifier: identifier,
		Role:       role,
		Kind:       kind,
		At:         at,
		Text:       text,
	}
}
