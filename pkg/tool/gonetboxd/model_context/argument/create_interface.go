package argument

type CreateInterface struct {
	Device          string `json:"device"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	PhysicalAddress string `json:"physical_address"`
}
