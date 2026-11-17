package argument

type CreateCable struct {
	DeviceA    string `json:"device_a"`
	InterfaceA string `json:"interface_a"`
	DeviceB    string `json:"device_b"`
	InterfaceB string `json:"interface_b"`
}
