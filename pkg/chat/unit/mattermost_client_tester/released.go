package mattermost_client_tester

func (t *Tester) Released() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.released
}
