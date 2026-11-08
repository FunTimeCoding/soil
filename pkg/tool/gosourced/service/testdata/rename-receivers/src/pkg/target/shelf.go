package target

type Product struct {
	Name string
	y    int
}

type Shelf struct {
	product *Product
	s       int
}

func (y *Shelf) Name() string {
	return y.product.Name
}

func (y *Shelf) Count() int {
	return y.s
}

func (y *Shelf) Depth() int {
	return y.product.y
}

func (*Shelf) Kind() string {
	return "shelf"
}

func (y Shelf) Copy() Shelf {
	return y
}

type Other struct{}

func (y *Other) Name() string {
	return "other"
}

func Plain(y int) int {
	return y + 1
}
