package argument

type VirtualTunnelTermination struct {
	VirtualMachine string `json:"virtual_machine"`
	Tunnel         string `json:"tunnel"`
	Interface      string `json:"interface"`
	Role           string `json:"role"`
}
