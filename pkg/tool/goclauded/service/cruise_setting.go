package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"

func (s *Service) CruiseSetting(
	sessionIdentifier string,
) (*cruise_setting.Setting, error) {
	return s.store.CruiseSetting(sessionIdentifier)
}
