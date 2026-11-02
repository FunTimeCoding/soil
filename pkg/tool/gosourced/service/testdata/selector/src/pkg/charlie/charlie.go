package charlie

import "example/pkg/bravo"

func Charlie(b *bravo.Bravo) int {
	x := b.Run(1, 2)

	return x
}
