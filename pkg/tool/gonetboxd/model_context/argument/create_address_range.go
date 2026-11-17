package argument

type CreateAddressRange struct {
	Start       string `json:"start"`
	End         string `json:"end"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
