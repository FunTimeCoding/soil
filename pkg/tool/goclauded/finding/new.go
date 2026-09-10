package finding

func New(
	kind string,
	subject string,
	detail string,
	count int64,
) *Finding {
	return &Finding{
		Kind:    kind,
		Subject: subject,
		Detail:  detail,
		Count:   count,
	}
}
