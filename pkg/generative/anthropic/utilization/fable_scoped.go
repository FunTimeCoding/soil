package utilization

import "github.com/funtimecoding/soil/pkg/generative/constant"

func fableScoped(s *Scope) bool {
	return s != nil &&
		s.Model != nil &&
		s.Model.DisplayName == constant.UsageScopeFable
}
