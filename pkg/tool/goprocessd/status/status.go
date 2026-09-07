package status

type Status struct {
	Name      string `json:"name"`
	Command   string `json:"command"`
	Running   bool   `json:"running"`
	StartedAt string `json:"started_at,omitzero"`
}
