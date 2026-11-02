package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	"time"
)

func (c *Client) SetSilence(
	alert string,
	comment string,
	d time.Duration,
	r *rule_list.List,
) (string, error) {
	if alert == "" {
		return "", validation.New("alert is required")
	}

	if r.Find(alert) == nil {
		return "", validation.New("rule not found: %s", alert)
	}

	t := time.Now()

	if d == 0 {
		d = constant.AlertmanagerDefaultDuration
	}

	s, f := c.SilenceByRule(alert)

	if f != nil {
		if not_found.Is(f) {
			return c.PostSilence("", alert, comment, t, t.Add(d))
		}

		return "", f
	}

	return c.PostSilence(s.Identifier, alert, comment, *s.Start, t.Add(d))
}
