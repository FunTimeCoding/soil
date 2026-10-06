package box

type Box struct{}

func (b *Box) Size() int {
	return 1
}

func (b *Box) Weight() int {
	return 2
}

func (b *Box) String() string {
	return "box"
}
