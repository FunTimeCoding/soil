package item

type Face interface {
	Describe() string
}

type Item struct {
	name string
}

func (b *Item) Name() string {
	return b.name
}

func (b *Item) Describe() string {
	return "item"
}

func (Item) Error() string {
	return "item"
}

func (b *Item) Weight() int {
	return 1
}

func (*Item) Height() int {
	return 2
}

func (_ *Item) Width() int {
	return 3
}

func (b *Item) Shadow() int {
	for b := 0; b < 4; b++ {
		return b
	}

	return 0
}
