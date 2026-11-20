package system_result

func NewDetail(
	s *Summary,
	findings []*Finding,
) *Detail {
	return &Detail{Summary: s, Findings: findings}
}
