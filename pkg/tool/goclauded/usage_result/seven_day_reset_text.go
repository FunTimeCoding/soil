package usage_result

func (r *Result) SevenDayResetText() string {
	return r.SevenDayReset.Local().Format("Mon 15:04")
}
