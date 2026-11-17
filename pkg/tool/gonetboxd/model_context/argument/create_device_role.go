package argument

type CreateDeviceRole struct {
	Name               string `json:"name"`
	VirtualMachineRole bool   `json:"virtual_machine_role"`
}
