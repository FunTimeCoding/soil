package remediation

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"

func rank(i *remediation_item.Item) int {
	if i.Addressed {
		return 1
	}

	return 0
}
