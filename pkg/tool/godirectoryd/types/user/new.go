package user

func New(
	account string,
	name string,
	surname string,
	mail string,
	unique string,
	distinguishedName string,
) *User {
	return &User{
		Account:           account,
		Name:              name,
		Surname:           surname,
		Mail:              mail,
		Unique:            unique,
		DistinguishedName: distinguishedName,
	}
}
