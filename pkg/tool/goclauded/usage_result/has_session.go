package usage_result

func (r *Result) HasSession() bool {
	return !r.FiveHourReset.IsZero()
}
