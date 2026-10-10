package summary_option

import "github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"

func New() *Option {
	return &Option{GroupBy: constant.Tool}
}
