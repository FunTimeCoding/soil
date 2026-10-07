package match

func New(
	identifier string,
	name string,
	alias string,
	field string,
) *Match {
	return &Match{
		Identifier: identifier,
		Name:       name,
		Alias:      alias,
		Field:      field,
	}
}
