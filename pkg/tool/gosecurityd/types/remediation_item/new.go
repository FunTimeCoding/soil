package remediation_item

func New(
	subject string,
	kind string,
	release string,
	source string,
) *Item {
	return &Item{
		Subject: subject,
		Kind:    kind,
		Release: release,
		Source:  source,
	}
}
