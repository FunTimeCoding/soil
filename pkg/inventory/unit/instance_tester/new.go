package instance_tester

func New(
	name string,
	host string,
) *Instance {
	return &Instance{Name: name, Host: host}
}
