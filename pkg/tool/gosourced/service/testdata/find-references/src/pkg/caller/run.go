package caller

import "example/pkg/target"

func Run() string {
	t := &target.Thing{}

	return target.Used + t.Ping()
}
