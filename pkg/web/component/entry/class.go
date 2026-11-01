package entry

func (e *Entry) Class(class string) *Entry {
	e.class = class

	return e
}
