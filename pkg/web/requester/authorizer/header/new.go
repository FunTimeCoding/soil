package header

func New(
	name string,
	value string,
) *Header {
	return &Header{name: name, value: value}
}
