package entry

func (e *Entry) ButtonClass(class string) *Entry {
	e.buttonClass = class

	return e
}
