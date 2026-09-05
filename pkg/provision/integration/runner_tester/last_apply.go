package runner_tester

func (o *Tester) LastApply() *ApplyCall {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.applied[len(o.applied)-1]
}
