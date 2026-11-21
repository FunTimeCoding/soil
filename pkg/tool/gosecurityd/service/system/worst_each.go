package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"slices"
)

func worstEach(
	findings []*system_result.Finding,
) map[string]*system_result.Finding {
	result := map[string]*system_result.Finding{}

	for _, f := range findings {
		current, okay := result[f.Vulnerability]

		if !okay ||
			slices.Index(constant.KindOrder, f.Kind()) >
				slices.Index(constant.KindOrder, current.Kind()) {
			result[f.Vulnerability] = f
		}
	}

	return result
}
