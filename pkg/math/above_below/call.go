package above_below

import "github.com/funtimecoding/soil/pkg/math/numeric"

func Call[T numeric.Number](
	value T,
	magnitude T,
	above func(),
	below func(),
) {
	if value > magnitude {
		above()
	} else if value*-1 > magnitude {
		below()
	}
}
