package basic

func New(
	user string,
	password string,
) *Basic {
	return &Basic{user: user, password: password}
}
