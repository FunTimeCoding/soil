package argument

type SetInterfacePhysicalAddress struct {
	Device    string `json:"device"`
	Interface string `json:"interface"`
	Address   string `json:"address"`
}
