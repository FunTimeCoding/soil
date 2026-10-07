package mattermost_client_tester

func (t *Tester) Refuse(count int) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.refusals = count
}
