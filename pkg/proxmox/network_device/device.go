package network_device

type Device struct {
	Name            string
	Model           string
	Interface       string
	HardwareAddress string
	Bridge          string
	Vlan            int
}
