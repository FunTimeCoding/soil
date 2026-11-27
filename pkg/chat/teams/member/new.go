package member

func New(
	name string,
	user string,
) *Member {
	return &Member{Name: name, User: user}
}
