package silence_tester

import "github.com/prometheus/alertmanager/api/v2/models"

func NewMatcher(
	name string,
	value string,
	equal bool,
	regex bool,
) *models.Matcher {
	return &models.Matcher{
		Name:    &name,
		Value:   &value,
		IsEqual: &equal,
		IsRegex: &regex,
	}
}
