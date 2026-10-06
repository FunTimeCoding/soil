package shape

type Shape struct {
	Width int
}

func New() *Shape {
	return &Shape{Width: 1}
}
