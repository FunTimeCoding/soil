package echo

import "example/delta"

func Pool() delta.Servers {
	return delta.Servers{{Port: 1}, {Port: 2}}
}
