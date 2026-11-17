package argument

type DeviceTunnelTermination struct {
	Device    string `json:"device"`
	Tunnel    string `json:"tunnel"`
	Interface string `json:"interface"`
	Role      string `json:"role"`
}
