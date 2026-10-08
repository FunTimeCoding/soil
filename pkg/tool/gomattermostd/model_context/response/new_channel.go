package response

func NewChannel(
	identifier string,
	name string,
	displayName string,
	kind string,
	purpose string,
	header string,
) *Channel {
	return &Channel{
		Identifier:  identifier,
		Name:        name,
		DisplayName: displayName,
		Type:        kind,
		Purpose:     purpose,
		Header:      header,
	}
}
