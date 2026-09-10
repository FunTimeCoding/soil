package receipt

func New(identifier string, name string) *Receipt {
	return &Receipt{Identifier: identifier, Name: name}
}
