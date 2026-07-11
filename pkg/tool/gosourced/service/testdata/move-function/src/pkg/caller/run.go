package caller

import "example/pkg/target"

func Run() bool {
	return target.IsValid("alfa")
}
