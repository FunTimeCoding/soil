package mattermost_client_tester

func (t *Tester) Accepted() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.accepted
}
