package usage_result

func (r *Result) HasFable() bool {
	return r.FableReset != ""
}
