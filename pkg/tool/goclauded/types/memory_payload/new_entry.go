package memory_payload

func NewEntry(
	identifier int64,
	name string,
) *Entry {
	return &Entry{Identifier: identifier, Name: name}
}
