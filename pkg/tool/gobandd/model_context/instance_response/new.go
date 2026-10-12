package instance_response

func New(
	name string,
	host string,
	active bool,
) *Response {
	return &Response{Name: name, Host: host, Active: active}
}
