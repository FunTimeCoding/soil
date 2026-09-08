package usage_result

func (r *Result) FableResetText() string {
	if r.FableResetAt.IsZero() {
		return r.FableReset
	}

	return r.FableResetAt.Local().Format("Mon 15:04")
}
