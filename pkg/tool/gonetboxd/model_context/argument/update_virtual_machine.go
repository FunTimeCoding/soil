package argument

type UpdateVirtualMachine struct {
	VirtualMachine string          `json:"virtual_machine"`
	Name           string          `json:"name"`
	PrimaryAddress string          `json:"primary_address"`
	Platform       string          `json:"platform"`
	Role           string          `json:"role"`
	Tenant         string          `json:"tenant"`
	Cores          float64         `json:"cores"`
	Memory         float64         `json:"memory"`
	Status         string          `json:"status"`
	CustomFields   map[string]any  `json:"custom_fields"`
	LocalContext   *map[string]any `json:"local_context"`
}
