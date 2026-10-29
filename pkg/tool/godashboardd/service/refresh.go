package service

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"

func (s *Service) Refresh() []error {
	result := map[string][]string{}
	var failed []error

	for _, v := range s.board.Entries() {
		var values []string
		var e []error

		switch {
		case v.Widget == constant.NextcloudWidget:
			values, e = one(s.usageValues())
		case v.Widget == constant.ArgocdWidget:
			values, e = one(s.argocdValues())
		case len(v.Rows) == 0:
			continue
		default:
			values, e = s.rowValues(v.Rows)
		}

		result[v.Label] = values
		failed = append(failed, e...)
	}

	s.mutex.Lock()
	s.values = result
	s.mutex.Unlock()
	s.notifier.Notify()

	return failed
}
