package relay

import "example/pkg/target"

func Size() int {
	return target.Size() + 1
}
