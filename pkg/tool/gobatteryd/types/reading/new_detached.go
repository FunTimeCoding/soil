package reading

func NewDetached(
	server string,
	battery string,
	driverConnected bool,
) *Reading {
	return &Reading{
		Server:          server,
		Battery:         battery,
		DriverConnected: driverConnected,
	}
}
