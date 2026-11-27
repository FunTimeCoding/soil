package user

func New(
	identifier string,
	name string,
	mail string,
) *User {
	return &User{Identifier: identifier, Name: name, Mail: mail}
}
