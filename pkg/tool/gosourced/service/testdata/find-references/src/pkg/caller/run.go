package caller

import (
	"example/pkg/target"
	"example/pkg/target/constant"
)

func Run() string {
	t := &target.Thing{}

	return constant.Used + t.Ping()
}
