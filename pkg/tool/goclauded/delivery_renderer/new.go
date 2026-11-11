package delivery_renderer

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"time"
)

func New(
	messages map[uint]*message.Message,
	location *time.Location,
) *Renderer {
	return &Renderer{messages: messages, location: location}
}
