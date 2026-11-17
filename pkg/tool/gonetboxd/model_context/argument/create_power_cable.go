package argument

type CreatePowerCable struct {
	OutletDevice string `json:"outlet_device"`
	Outlet       string `json:"outlet"`
	PortDevice   string `json:"port_device"`
	Port         string `json:"port"`
}
