package profile

func NewHeading(
	identifier int64,
	name string,
) *Heading {
	return &Heading{Identifier: identifier, Name: name}
}
