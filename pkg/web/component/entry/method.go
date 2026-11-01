package entry

func (e *Entry) method() string {
	if e.query {
		return "get"
	}

	return "post"
}
