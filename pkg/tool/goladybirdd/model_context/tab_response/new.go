package tab_response

func New(
	identifier string,
	locator string,
) *Response {
	return &Response{Identifier: identifier, Locator: locator}
}
