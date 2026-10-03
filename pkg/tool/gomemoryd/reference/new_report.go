package reference

func NewReport(
	identifier int64,
	name string,
	findings []*Finding,
) *Report {
	return &Report{Identifier: identifier, Name: name, Findings: findings}
}
