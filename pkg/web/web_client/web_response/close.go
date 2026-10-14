package web_response

import "github.com/funtimecoding/soil/pkg/errors"

func (r *Response) Close() {
	if r.Response != nil {
		errors.LogClose(r.Response.Body)
	}
}
