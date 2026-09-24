package entry

func (e *Entry) IsNew() bool {
	return e.Previous == nil
}
