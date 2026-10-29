package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) WebSocket() *model.WebSocketClient {
	c.socketMutex.Lock()
	defer c.socketMutex.Unlock()

	return c.webSocket
}
