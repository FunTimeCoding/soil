package target

type Size struct {
	Width  int
	Height int
}

func NewSize(
	width int,
	height int,
) *Size {
	return &Size{Width: width, Height: height}
}

func add(
	a int,
	b int,
) int {
	return a + b
}

func small(
	width int,
	height int,
) *Size {
	return NewSize(width, height)
}

func plain(
	a int,
	b int,
) int {
	return add(a, b)
}

func reordered(
	a int,
	b int,
) int {
	return add(b, a)
}

func wrapped(
	a int,
	b int,
) int {
	return add(a, b) + 1
}

func multiple(
	a int,
	b int,
) int {
	c := add(a, b)

	return c
}

func bound(a int) int {
	return add(a, 1)
}

func none() *Size {
	return NewSize(1, 2)
}

func (s *Size) Scale(factor int) *Size {
	return NewSize(s.Width*factor, s.Height*factor)
}

func chained(factor int) *Size {
	return NewSize(1, 2).Scale(factor)
}
