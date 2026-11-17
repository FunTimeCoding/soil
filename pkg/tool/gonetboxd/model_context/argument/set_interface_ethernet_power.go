package argument

type SetInterfaceEthernetPower struct {
	Device    string `json:"device"`
	Interface string `json:"interface"`
	Mode      string `json:"mode"`
	Type      string `json:"type"`
}
