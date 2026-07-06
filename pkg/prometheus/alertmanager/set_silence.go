package alertmanager

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"time"
)

func (c *Client) SetSilence(
	alert string,
	comment string,
	d time.Duration,
) (string, error) {
	if alert == "" {
		return "", validation.New("alert is required")
	}

	exists, e := c.RuleExists(alert)

	if e != nil {
		return "", e
	}

	if !exists {
		return "", validation.New("rule not found: %s", alert)
	}

	t := time.Now()

	if d == 0 {
		d = constant.DefaultDuration
	}

	s, f := c.SilenceByRule(alert)

	if f != nil {
		if errors.Is(f, constant.ErrorNotFound) {
			return c.PostSilence("", alert, comment, t, t.Add(d))
		}

		return "", f
	}

	return c.PostSilence(
		s.Identifier,
		alert,
		comment,
		*s.Start,
		t.Add(d),
	)
}
