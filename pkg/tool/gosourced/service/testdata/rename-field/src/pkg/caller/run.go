package caller

import "example/pkg/target"

func Run() string {
	s := &target.Store{Name: "alfa"}

	return s.Name
}
