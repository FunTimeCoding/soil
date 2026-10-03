package rule_list

import "github.com/prometheus/common/model"

func LabelsSame(
	a model.LabelSet,
	b model.LabelSet,
) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if b[k] != v {
			return false
		}
	}

	return true
}
