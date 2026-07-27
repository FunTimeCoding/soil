package scan

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"strings"
)

func (s *Service) checkSuffix(path string) {
	if !strings.HasSuffix(s.Name, "d") && s.hasCapability() {
		s.addConcern(
			constant.MissingSuffixKey,
			constant.MissingSuffixText,
			path,
		)
	}
}
