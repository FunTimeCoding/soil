package host

func New(
	name string,
	platform string,
	hardwareAddresses []string,
) *Host {
	return &Host{
		Name:              name,
		Platform:          platform,
		HardwareAddresses: hardwareAddresses,
	}
}
