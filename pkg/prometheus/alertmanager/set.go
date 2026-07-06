package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_enricher"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/field_changer"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/name_filter"
)

func (c *Client) Set(
	e *alert_enricher.Enricher,
	h *field_changer.Changer,
	n *name_filter.Filter,
	l *label_filter.Filter,
) *Client {
	c.enricher = e
	c.changer = h
	c.nameFilter = n
	c.labelFilter = l

	return c
}
