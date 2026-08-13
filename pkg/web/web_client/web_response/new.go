package web_response

import "net/http"

func New(
	r *http.Response,
	milliseconds int64,
) *Response {
	return &Response{Response: r, Milliseconds: milliseconds}
}
