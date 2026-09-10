package caller

import "example/pkg/target"

func Run() int {
	return target.Compute()
}
