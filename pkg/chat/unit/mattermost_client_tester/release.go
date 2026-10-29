package mattermost_client_tester

func (t *Tester) release() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.released++
}
