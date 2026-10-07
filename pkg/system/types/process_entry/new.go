package process_entry

func New(
	identifier int32,
	parent int32,
	name string,
) *Entry {
	return &Entry{Identifier: identifier, Parent: parent, Name: name}
}
