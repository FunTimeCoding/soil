package path_move

func New(
	from string,
	to string,
) *Move {
	return &Move{Old: from, New: to}
}
