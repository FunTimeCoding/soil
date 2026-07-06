package opsgenie

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert/alert_processor"

func (c *Client) processor() *alert_processor.Processor {
	return alert_processor.New(c.enricher)
}
