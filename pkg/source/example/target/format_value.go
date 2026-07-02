package target

func (e *Entry) FormatValue() string {
	if IsValid(e.Value) {
		return e.Value
	}

	return "empty"
}
