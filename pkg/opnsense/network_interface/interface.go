package network_interface

type Interface struct {
	Device     string
	Status     string
	Media      string
	MacAddress string
	Mtu        string
	Addresses  []string
}
