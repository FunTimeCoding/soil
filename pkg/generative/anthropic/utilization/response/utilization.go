package response

type Utilization struct {
	Limits []Limit `json:"limits"`
}
