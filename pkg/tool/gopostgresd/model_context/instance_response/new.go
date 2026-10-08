package instance_response

func New(
	name string,
	host string,
	port int,
	database string,
	active bool,
) *Response {
	return &Response{
		Name:     name,
		Host:     host,
		Port:     port,
		Database: database,
		Active:   active,
	}
}
