package holiday

func NewEaster(offset int) *Holiday {
	return &Holiday{movable: true, offset: offset}
}
