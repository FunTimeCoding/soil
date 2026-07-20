package service

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"

func (s *Service) Refresh() {
	result := map[string][]string{}

	for _, v := range s.board.Entries() {
		if v.Widget == constant.NextcloudWidget {
			result[v.Label] = s.usageValues()

			continue
		}

		if v.Widget == constant.ArgocdWidget {
			result[v.Label] = s.argocdValues()

			continue
		}

		if len(v.Rows) == 0 {
			continue
		}

		result[v.Label] = s.rowValues(v.Rows)
	}

	s.mutex.Lock()
	s.values = result
	s.mutex.Unlock()
	s.notifier.Notify()
}
