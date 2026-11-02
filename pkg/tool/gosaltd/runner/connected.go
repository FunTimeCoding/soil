package runner

func (r *Runner) Connected() bool {
	select {
	case <-r.connected:
		return true
	default:
		return false
	}
}
