package mattermost_client_tester

func (t *Tester) Stop() {
	t.server.CloseClientConnections()
	t.server.Close()
}
