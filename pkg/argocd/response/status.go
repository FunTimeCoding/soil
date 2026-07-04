package response

type Status struct {
	Sync   Sync   `json:"sync"`
	Health Health `json:"health"`
}
