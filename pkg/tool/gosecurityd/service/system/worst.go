package system

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"

func worst(
	findings []*system_result.Finding,
	name string,
) *system_result.Finding {
	return worstEach(findings)[name]
}
