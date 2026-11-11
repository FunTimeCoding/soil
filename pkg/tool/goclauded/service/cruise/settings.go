package cruise

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"

func (s *Service) Settings(
	sessionIdentifiers []string,
) (map[string]*cruise_setting.Setting, error) {
	return s.store.CruiseSettings(sessionIdentifiers)
}
