package rule

import "github.com/prometheus/client_golang/api/prometheus/v1"

func NewRecord(
	u *v1.RecordingRule,
	g *v1.RuleGroup,
) *Rule {
	r := &Rule{Group: g.Name, Name: u.Name, RawRecord: u, RawGroup: g}

	return r.parse()
}
