package mutation

type Report struct {
	Module string  `json:"go_module"`
	Files  []*File `json:"files"`
}
