package mattermost_client_tester

func (t *Tester) refused() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if t.refusals == 0 {
		return false
	}

	t.refusals--

	return true
}
