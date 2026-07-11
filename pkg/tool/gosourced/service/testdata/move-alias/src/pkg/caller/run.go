package caller

import (
	"example/pkg/caller/constant"
	"example/pkg/target"
)

func Run() string {
	return constant.Local + target.Mode
}
