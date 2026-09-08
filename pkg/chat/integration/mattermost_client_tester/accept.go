package mattermost_client_tester

import "github.com/gorilla/websocket"

func (t *Tester) accept(c *websocket.Conn) {
	t.mutex.Lock()
	t.connection = c
	t.accepted++
	t.mutex.Unlock()
	t.once.Do(func() { close(t.ready) })
}
