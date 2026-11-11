package send_result

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func New(
	requested constant.Delivery,
	d constant.Delivery,
) *Result {
	return &Result{Requested: requested, Delivery: d}
}
