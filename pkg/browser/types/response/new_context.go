package response

func NewContext(
	identifier string,
	locator string,
) *Context {
	return &Context{Identifier: identifier, Locator: locator}
}
