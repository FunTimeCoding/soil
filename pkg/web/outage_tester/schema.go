package outage_tester

type schema struct {
	Required   []string            `json:"required"`
	Properties map[string]property `json:"properties"`
}
