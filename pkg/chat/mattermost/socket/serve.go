package socket

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
)

func (s *Socket) Serve(x context.Context) error {
	w := s.client.WebSocket()

	for {
		select {
		case <-x.Done():
			return nil
		case v, okay := <-w.EventChannel:
			if !okay {
				return dropped.Format(
					"%s",
					constant.MattermostEventChannelClosed,
				)
			}

			s.dispatch(v)
		case <-w.PingTimeoutChannel:
			return dropped.Format("%s", constant.MattermostPingTimeout)
		}
	}
}
