package response

type ResourceSummary struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
	Kind      string `json:"kind,omitempty"`
	Status    string `json:"status,omitempty"`
	Restarts  *int64 `json:"restarts,omitempty"`
	Age       string `json:"age,omitempty"`
}
