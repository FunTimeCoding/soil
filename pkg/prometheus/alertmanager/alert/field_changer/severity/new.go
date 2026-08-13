package severity

func New(
	alert string,
	from string,
	to string,
) *Severity {
	return &Severity{Alert: alert, From: from, To: to}
}
