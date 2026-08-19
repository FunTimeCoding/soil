package argument

type SignRequest struct {
	Authority string  `json:"authority"`
	Kind      string  `json:"kind"`
	Request   string  `json:"request"`
	ValidDay  float64 `json:"valid_day"`
}
