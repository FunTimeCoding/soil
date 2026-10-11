package delivery

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"time"
)

func New(
	messages map[uint]*message.Message,
	location *time.Location,
) *Delivery {
	return &Delivery{messages: messages, location: location}
}
