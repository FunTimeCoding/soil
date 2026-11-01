package socket

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
)

func (s *Socket) Serve(x context.Context) error {
	w := s.client.WebSocket()

	for {
		select {
		case <-x.Done():
			return nil
		case v, okay := <-w.EventChannel:
			if !okay {
				return dropped.Format("%s", constant.EventChannelClosed)
			}

			s.dispatch(v)
		case <-w.PingTimeoutChannel:
			return dropped.Format("%s", constant.PingTimeout)
		}
	}
}
