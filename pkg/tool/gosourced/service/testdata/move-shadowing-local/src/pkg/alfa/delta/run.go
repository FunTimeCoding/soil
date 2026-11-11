package delta

import "example/pkg/alfa/bravo"

func Run() int {
	echo := bravo.Limit()

	return echo * 2
}
