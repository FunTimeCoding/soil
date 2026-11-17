package argument

type DeviceCustomField struct {
	Device string   `json:"device"`
	Field  string   `json:"field"`
	Values []string `json:"values"`
}
