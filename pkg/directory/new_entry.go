package directory

func NewEntry(
	unique string,
	account string,
	mail string,
	name string,
	distinguishedName string,
) *Entry {
	return &Entry{
		Unique:            unique,
		Account:           account,
		Mail:              mail,
		Name:              name,
		DistinguishedName: distinguishedName,
	}
}
