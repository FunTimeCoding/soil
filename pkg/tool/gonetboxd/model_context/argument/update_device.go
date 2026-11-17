package argument

type UpdateDevice struct {
	Device         string          `json:"device"`
	Name           string          `json:"name"`
	PrimaryAddress string          `json:"primary_address"`
	Location       string          `json:"location"`
	Platform       string          `json:"platform"`
	Type           string          `json:"type"`
	Role           string          `json:"role"`
	Site           string          `json:"site"`
	Tenant         string          `json:"tenant"`
	Status         string          `json:"status"`
	Serial         string          `json:"serial"`
	Description    string          `json:"description"`
	CustomFields   map[string]any  `json:"custom_fields"`
	LocalContext   *map[string]any `json:"local_context"`
}
