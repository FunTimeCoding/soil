package transcript

type Line struct {
	Cwd     string  `json:"cwd"`
	Message Message `json:"message"`
}
