package transcript

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/permission"
)

func (s *Service) Permissions(days int) (*permission.Report, error) {
	if days != constant.PermissionWindowDays {
		return s.computePermissions(days)
	}

	s.reports.RLock()
	held := s.heldPermissions
	s.reports.RUnlock()

	if held != nil {
		return held, nil
	}

	result, e := s.computePermissions(days)

	if e != nil {
		return nil, e
	}

	s.reports.Lock()
	s.heldPermissions = result
	s.reports.Unlock()

	return result, nil
}
