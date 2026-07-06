package rule_list

import "github.com/funtimecoding/soil/pkg/prometheus/rule"

func (l *List) Record() []*rule.Rule {
	var rules []*rule.Rule

	for _, r := range l.rules {
		if r.IsRecord() {
			rules = append(rules, r)
		}
	}

	return rules
}
