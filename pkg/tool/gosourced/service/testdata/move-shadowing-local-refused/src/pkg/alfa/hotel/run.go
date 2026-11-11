package hotel

import "example/pkg/alfa/bravo"

func Run(
	e int,
	c int,
	h int,
	o int,
) int {
	echo := bravo.Limit()

	return echo * e * c * h * o
}
