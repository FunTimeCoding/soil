package process

func NewEntry(
	identifier int32,
	parent int32,
	name string,
) *Entry {
	return &Entry{Identifier: identifier, Parent: parent, Name: name}
}
