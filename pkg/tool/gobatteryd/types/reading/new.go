package reading

func New(
	server string,
	battery string,
	variables map[string]string,
) *Reading {
	return &Reading{
		Server:          server,
		Battery:         battery,
		Attached:        true,
		DriverConnected: true,
		Variables:       variables,
	}
}
