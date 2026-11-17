package argument

type CreateVirtualAddress struct {
	VirtualMachine string `json:"virtual_machine"`
	Interface      string `json:"interface"`
	Address        string `json:"address"`
	Status         string `json:"status"`
}
