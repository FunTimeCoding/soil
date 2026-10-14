package runner_tester

func (o *Tester) ApplyCount() int {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return len(o.applied)
}
