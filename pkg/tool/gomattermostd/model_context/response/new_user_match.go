package response

func NewUserMatch(
	identifier string,
	username string,
	firstName string,
	lastName string,
	nickname string,
	email string,
) *UserMatch {
	return &UserMatch{
		Identifier: identifier,
		Username:   username,
		FirstName:  firstName,
		LastName:   lastName,
		Nickname:   nickname,
		Email:      email,
	}
}
