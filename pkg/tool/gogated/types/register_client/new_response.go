package register_client

func NewResponse(
	clientIdentifier string,
	clientSecret string,
	redirectLocators []string,
) *Response {
	return &Response{
		ClientIdentifier: clientIdentifier,
		ClientSecret:     clientSecret,
		RedirectLocators: redirectLocators,
	}
}
