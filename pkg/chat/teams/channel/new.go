package channel

func New(
	identifier string,
	name string,
) *Channel {
	return &Channel{Identifier: identifier, Name: name}
}
