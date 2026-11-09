package shape

type Solid interface {
	Sizer
	Weight() int
}
