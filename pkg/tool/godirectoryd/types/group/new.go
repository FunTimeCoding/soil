package group

func New(
	name string,
	number int,
	member []string,
	distinguishedName string,
) *Group {
	return &Group{
		Name:              name,
		Number:            number,
		Member:            member,
		DistinguishedName: distinguishedName,
	}
}
