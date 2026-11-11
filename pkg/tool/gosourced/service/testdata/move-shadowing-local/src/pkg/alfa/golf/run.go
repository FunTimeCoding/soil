package golf

import "example/pkg/alfa/bravo"

func Run(e int) int {
	echo := bravo.Limit()

	return echo * e
}
