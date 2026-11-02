package door

func New(
	identifier int64,
	name string,
	relation string,
	source string,
) *Door {
	return &Door{
		Identifier: identifier,
		Name:       name,
		Relation:   relation,
		Source:     source,
	}
}
