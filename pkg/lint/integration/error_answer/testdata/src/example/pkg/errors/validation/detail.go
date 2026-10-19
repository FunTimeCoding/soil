package validation

type Detail struct {
	Message string
}

func (d *Detail) Error() string {
	return d.Message
}
