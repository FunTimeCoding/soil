package system_selection

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"maps"
	"net/url"
	"slices"
	"strconv"
)

func (s *Selection) Query(fallback float64) url.Values {
	result := url.Values{}

	if s.Show != "" {
		result.Set(constant.ShowField, s.Show)
	}

	if s.Minimum != fallback {
		result.Set(
			constant.ScoreField,
			strconv.FormatFloat(s.Minimum, 'f', -1, 64),
		)
	}

	for _, key := range slices.Sorted(maps.Keys(s.Labels)) {
		result.Add(
			constant.LabelField,
			join.Empty(key, constant.LabelSeparator, s.Labels[key]),
		)
	}

	if s.Search != "" {
		result.Set(constant.SearchField, s.Search)
	}

	return result
}
