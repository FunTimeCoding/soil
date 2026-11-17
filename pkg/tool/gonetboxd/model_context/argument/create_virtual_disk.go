package argument

type CreateVirtualDisk struct {
	VirtualMachine string  `json:"virtual_machine"`
	Name           string  `json:"name"`
	Size           float64 `json:"size"`
}
