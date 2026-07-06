package rule_list

import "github.com/funtimecoding/soil/pkg/prometheus/rule"

func (l *List) Alert() []*rule.Rule {
	var rules []*rule.Rule

	for _, r := range l.rules {
		if r.IsAlert() {
			rules = append(rules, r)
		}
	}

	return rules
}
