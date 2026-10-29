package selector

func New(
	identifier string,
	title string,
	locator string,
) *Selector {
	return &Selector{Identifier: identifier, Title: title, Locator: locator}
}
