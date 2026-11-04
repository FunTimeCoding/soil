package notify_call

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

type Call struct {
	Callsign string
	Source   string
	Body     string
	Delivery constant.Delivery
}
