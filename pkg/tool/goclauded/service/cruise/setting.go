package cruise

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"

func (s *Service) Setting(
	sessionIdentifier string,
) (*cruise_setting.Setting, error) {
	return s.store.CruiseSetting(sessionIdentifier)
}
