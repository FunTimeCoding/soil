package mock_directory

func (d *Directory) SetMail(
	unique string,
	mail string,
) {
	for _, entry := range d.entries {
		if entry.Unique == unique {
			entry.Mail = mail
		}
	}
}
