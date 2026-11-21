package system_selection

import "maps"

func (s *Selection) clone() *Selection {
	result := *s
	result.Labels = maps.Clone(s.Labels)

	if result.Labels == nil {
		result.Labels = map[string]string{}
	}

	return &result
}
