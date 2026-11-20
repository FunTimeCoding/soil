package package_result

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"

func New(s *system_result.Summary) *Result {
	return &Result{Summary: s}
}
