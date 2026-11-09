package alert

func (a *Alert) descriptionName(s string) string {
	if a.descriptionToName != nil {
		return a.descriptionToName(s)
	}

	return s
}
