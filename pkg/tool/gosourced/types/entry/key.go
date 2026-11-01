package entry

func (e *Entry) Key() string {
	if e.To != nil {
		return e.To.Key()
	}

	return e.From.Key()
}
