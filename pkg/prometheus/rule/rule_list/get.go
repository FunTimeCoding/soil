package rule_list

import "github.com/funtimecoding/soil/pkg/prometheus/rule"

func (l *List) Get() []*rule.Rule {
	return l.rules
}
