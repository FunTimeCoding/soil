package update_request

type Request struct {
	Status     string `json:"status,omitempty"`
	AssignedTo string `json:"assignedTo,omitempty"`
}
