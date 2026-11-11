package cruise

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"time"
)

func (s *Service) Grace(c *cruise_setting.Setting) time.Duration {
	if c.Grace != nil {
		return *c.Grace
	}

	return s.defaultGrace
}
