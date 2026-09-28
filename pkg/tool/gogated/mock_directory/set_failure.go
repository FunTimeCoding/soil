package mock_directory

func (d *Directory) SetFailure(e error) {
	d.failure = e
}
