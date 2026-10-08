package response

func NewJoinedChannel(
	identifier string,
	name string,
	displayName string,
	kind string,
	lastPostAt string,
) *JoinedChannel {
	return &JoinedChannel{
		Identifier:  identifier,
		Name:        name,
		DisplayName: displayName,
		Type:        kind,
		LastPostAt:  lastPostAt,
	}
}
