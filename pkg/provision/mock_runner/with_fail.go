package mock_runner

func (r *Runner) WithFail(e error) *Runner {
	r.fail = e

	return r
}
