package mattermost

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/gorilla/websocket"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) ping(
	s *model.WebSocketClient,
	stop chan struct{},
	t *time.Ticker,
) {
	defer t.Stop()

	for {
		select {
		case <-stop:
			return
		case <-t.C:
			e := s.Conn.WriteControl(
				websocket.PingMessage,
				nil,
				time.Now().Add(constant.MattermostPingWait),
			)

			if e != nil {
				return
			}
		}
	}
}
