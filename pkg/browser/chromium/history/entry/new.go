package entry

func New(
	title string,
	locator string,
) *Entry {
	return &Entry{Title: title, Locator: locator}
}
