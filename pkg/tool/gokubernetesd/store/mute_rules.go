package store

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store/mute_rule"

func (s *Store) MuteRules() ([]mute_rule.MuteRule, error) {
	var result []mute_rule.MuteRule

	return result, s.database.Order("created_at").Find(&result).Error
}
