package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store/mute_rule"
	"strings"
)

func (s *Store) IsMuted(
	reason string,
	message string,
	cluster string,
) (bool, error) {
	var rules []mute_rule.MuteRule
	e := s.database.Find(&rules).Error

	if e != nil {
		return false, e
	}

	for _, rule := range rules {
		if rule.Cluster != "" && rule.Cluster != cluster {
			continue
		}

		if rule.Reason != reason {
			continue
		}

		if rule.Message != "" && !strings.Contains(
			strings.ToLower(message),
			strings.ToLower(rule.Message),
		) {
			continue
		}

		return true, nil
	}

	return false, nil
}
