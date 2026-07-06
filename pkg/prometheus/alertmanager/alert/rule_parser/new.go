package rule_parser

import "github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"

func New(r *rule_list.List) *Parser {
	return &Parser{list: r}
}
