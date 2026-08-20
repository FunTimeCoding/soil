package response

type NetworkInterface struct {
	Device     string    `json:"device"`
	Status     string    `json:"status"`
	Media      string    `json:"media"`
	MacAddress string    `json:"macaddr"`
	Mtu        string    `json:"mtu"`
	Version4   []Address `json:"ipv4"`
	Version6   []Address `json:"ipv6"`
}
