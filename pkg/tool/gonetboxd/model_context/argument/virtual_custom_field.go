package argument

type VirtualCustomField struct {
	VirtualMachine string   `json:"virtual_machine"`
	Field          string   `json:"field"`
	Values         []string `json:"values"`
}
