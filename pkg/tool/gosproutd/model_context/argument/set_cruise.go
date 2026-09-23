package argument

type SetCruise struct {
	Session string  `json:"session"`
	Mode    string  `json:"mode"`
	Pace    float64 `json:"pace"`
}
