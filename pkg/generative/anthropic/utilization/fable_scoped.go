package utilization

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/response"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func fableScoped(s *response.Scope) bool {
	return s != nil &&
		s.Model != nil &&
		s.Model.DisplayName == constant.UsageScopeFable
}
