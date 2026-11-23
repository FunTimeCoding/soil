package team

func New(
	identifier string,
	name string,
) *Team {
	return &Team{Identifier: identifier, Name: name}
}
