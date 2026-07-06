package service

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store/mute_rule"

func (s *Service) MuteRules() ([]mute_rule.MuteRule, error) {
	return s.store.MuteRules()
}
