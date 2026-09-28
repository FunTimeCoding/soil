package mock_directory

func (d *Directory) SetMember(
	account string,
	member bool,
) {
	d.members[account] = member
}
