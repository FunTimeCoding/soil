package send_result

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (r *Result) Immediate() bool {
	return r.Delivery == constant.DeliveryImmediate && len(r.Held) == 0
}
