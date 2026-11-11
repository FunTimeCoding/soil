package india

import "example/pkg/alfa/bravo"

func Run() int {
	echo := bravo.Limit()

	for e := range 2 {
		echo += e
	}

	return echo
}
