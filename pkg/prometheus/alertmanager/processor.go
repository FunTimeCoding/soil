package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_processor"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/rule_parser"
)

func (c *Client) processor(p *advanced_option.Alert) (*alert_processor.Processor, error) {
	v, e := c.Rules()

	if e != nil {
		return nil, e
	}

	return alert_processor.New(
		p,
		c.enricher,
		c.changer,
		c.nameFilter,
		c.labelFilter,
		rule_parser.New(v),
	), nil
}
