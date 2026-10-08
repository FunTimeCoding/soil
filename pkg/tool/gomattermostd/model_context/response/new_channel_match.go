package response

func NewChannelMatch(
	identifier string,
	name string,
	displayName string,
	kind string,
) *ChannelMatch {
	return &ChannelMatch{
		Identifier:  identifier,
		Name:        name,
		DisplayName: displayName,
		Type:        kind,
	}
}
