package rule_parser

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"

func (p *Parser) Runbook(name string) string {
	if p.list == nil {
		return constant.None
	}

	result := p.list.Runbook(name)

	if result == "" {
		result = constant.None
	}

	return result
}
