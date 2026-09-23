package unit

import "github.com/prometheus/alertmanager/api/v2/models"

type silenceMatcherCase struct {
	name          string
	matcherName   string
	matcherValue  string
	labels        models.LabelSet
	expectedMatch bool
	description   string
}
