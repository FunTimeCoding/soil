package intent

type Plan struct {
	Hypervisor string    `json:"hypervisor"`
	Written    bool      `json:"written"`
	Changes    []*Change `json:"changes"`
	Unknown    []string  `json:"unknown"`
	Ambiguous  []string  `json:"ambiguous"`
}
