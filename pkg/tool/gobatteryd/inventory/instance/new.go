package instance

func New(
	name string,
	host string,
	port int,
) *Instance {
	return &Instance{Name: name, Host: host, Port: port}
}
