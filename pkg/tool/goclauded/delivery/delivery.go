package delivery

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"time"
)

type Delivery struct {
	messages map[uint]*message.Message
	location *time.Location
}
