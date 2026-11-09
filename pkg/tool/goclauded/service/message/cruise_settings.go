package message

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"

func (s *Service) CruiseSettings(
	sessionIdentifiers []string,
) (map[string]*cruise_setting.Setting, error) {
	return s.store.CruiseSettings(sessionIdentifiers)
}
