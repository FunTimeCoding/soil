package runner

func (r *Runner) stopped() bool {
	select {
	case <-r.stop:
		return true
	default:
		return false
	}
}
