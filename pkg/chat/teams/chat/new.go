package chat

func New(
	identifier string,
	kind string,
) *Chat {
	return &Chat{Identifier: identifier, Kind: kind}
}
