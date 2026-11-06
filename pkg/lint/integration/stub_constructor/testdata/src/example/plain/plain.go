package plain

type Option struct {
	Name string
}

func New() *Option {
	return &Option{}
}
