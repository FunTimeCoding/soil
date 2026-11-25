package model_context

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/prometheus/common/model"
	"time"
)

func parseSince(
	raw string,
	now time.Time,
) (*time.Time, error) {
	if raw == "" {
		return nil, nil
	}

	if d, e := model.ParseDuration(raw); e == nil {
		result := now.Add(-time.Duration(d))

		return &result, nil
	}

	result, e := time.Parse(time.RFC3339, raw)

	if e != nil {
		return nil, validation.New(
			"updated_after is neither a duration (7d, 36h) nor RFC3339: %s",
			raw,
		)
	}

	return &result, nil
}
