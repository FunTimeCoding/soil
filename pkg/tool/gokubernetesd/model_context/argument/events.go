package argument

type Events struct {
	Namespace    string `json:"namespace"`
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Limit        int    `json:"limit"`
	IncludeMuted bool   `json:"include_muted"`
}
