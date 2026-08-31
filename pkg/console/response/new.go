package response

func New(body string, status int) *Response {
	return &Response{Body: body, Status: status}
}
