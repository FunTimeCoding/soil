package response

func NewUser(
	identifier string,
	username string,
	firstName string,
	lastName string,
	nickname string,
	email string,
	isBot bool,
) *User {
	return &User{
		Identifier: identifier,
		Username:   username,
		FirstName:  firstName,
		LastName:   lastName,
		Nickname:   nickname,
		Email:      email,
		IsBot:      isBot,
	}
}
