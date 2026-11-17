package argument

type SetVirtualInterfacePhysicalAddress struct {
	VirtualMachine string `json:"virtual_machine"`
	Interface      string `json:"interface"`
	Address        string `json:"address"`
}
