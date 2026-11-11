package delivery_renderer

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"time"
)

type Renderer struct {
	messages map[uint]*message.Message
	location *time.Location
}
