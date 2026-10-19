package answer

func NewCaptured(status int) *Answer {
	return &Answer{Status: status, Captured: true}
}
