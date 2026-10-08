package utilization

type Response struct {
	Limits []Limit `json:"limits"`
}
