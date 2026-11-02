package response

type ArgocdSummary struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Sync      string `json:"sync"`
	Health    string `json:"health"`
	Operation string `json:"operation,omitempty"`
}
