package path_move

func New(
	from string,
	to string,
) *Move {
	return &Move{From: from, To: to}
}
