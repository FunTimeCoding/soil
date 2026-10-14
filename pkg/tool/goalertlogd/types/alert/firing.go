package alert

func (a *Alert) Firing() bool {
	return a.End == nil
}
