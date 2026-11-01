package mock_client

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
)

func (c *Client) Subscribe(_ string, _ []string, _ *logger.Logger, _ face.Reporter) (<-chan *stream_event.Event, func()) {
	return nil, func(){}
}
