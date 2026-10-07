package trigger

type Request struct {
	Parameters map[string]any
	Update     bool
	Response   chan *Result
}
