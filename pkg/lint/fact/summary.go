package fact

type Summary struct {
	Arranges  []string            `json:"arranges"`
	Delegates map[string][]string `json:"delegates"`
}
