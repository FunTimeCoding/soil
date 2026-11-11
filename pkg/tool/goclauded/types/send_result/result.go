package send_result

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

type Result struct {
	Requested constant.Delivery
	Delivery  constant.Delivery
	Held      []string
}
