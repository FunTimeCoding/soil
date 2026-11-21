package remediation_result

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"

func New() *Result {
	return &Result{Items: []*remediation_item.Item{}, Kinds: map[string]int{}}
}
