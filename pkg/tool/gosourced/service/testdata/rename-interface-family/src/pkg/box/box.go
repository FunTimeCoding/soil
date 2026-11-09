package box

type Box struct{}

func (b *Box) Size() int {
	return 1
}

func (b *Box) String() string {
	return "box"
}
