package remediation_result

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"

type Result struct {
	Items       []*remediation_item.Item
	Open        int
	Addressed   int
	Unimportant int
	Kinds       map[string]int
}
