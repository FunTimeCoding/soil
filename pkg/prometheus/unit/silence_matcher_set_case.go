package unit

import "github.com/prometheus/alertmanager/api/v2/models"

type silenceMatcherSetCase struct {
	name          string
	matchers      []*models.Matcher
	labels        models.LabelSet
	expectedMatch bool
	description   string
}
