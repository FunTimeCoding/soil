package service

import (
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"time"
)

func (s *Service) LabelValues(
	instance string,
	label string,
	matches []string,
	since time.Time,
) (*label_result.Result, error) {
	c, e := s.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.LabelValues(label, matches, since)
}
