package place

func New(
	kind string,
	identifier int32,
	name string,
) *Place {
	return &Place{Kind: kind, Identifier: identifier, Name: name}
}
