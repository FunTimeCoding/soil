package mattermost_client_tester

func (t *Tester) countPing() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.pings++
}
