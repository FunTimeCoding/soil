package pointer

func NewConcern(key string) *Resolution {
	return &Resolution{Concern: key}
}
