package receiver_rename

func New(
	typeName string,
	name string,
) *Rename {
	return &Rename{Type: typeName, Name: name}
}
