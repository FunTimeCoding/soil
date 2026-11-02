package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_processor"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/rule_parser"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
)

func (c *Client) processor(
	p *advanced_option.Alert,
	r *rule_list.List,
) *alert_processor.Processor {
	return alert_processor.New(
		p,
		c.enricher,
		c.changer,
		c.nameFilter,
		c.labelFilter,
		rule_parser.New(r),
	)
}
