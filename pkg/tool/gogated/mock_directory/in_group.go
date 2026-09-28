package mock_directory

func (d *Directory) InGroup(account string) (bool, error) {
	if d.failure != nil {
		return false, d.failure
	}

	return d.members[account], nil
}
