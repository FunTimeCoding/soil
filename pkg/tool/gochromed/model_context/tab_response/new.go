package tab_response

func New(
	identifier string,
	title string,
	locator string,
) *Response {
	return &Response{
		Identifier: identifier,
		Title:      title,
		Locator:    locator,
	}
}
