package mattermost_client_tester

func (t *Tester) Pings() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.pings
}
