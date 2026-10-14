package cruise

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func New() *Reading {
	return &Reading{Mode: constant.CruiseOff}
}
