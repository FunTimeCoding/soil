package pointer

func NewLive(target string) *Resolution {
	return &Resolution{Target: target}
}
