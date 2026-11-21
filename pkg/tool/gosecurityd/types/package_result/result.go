package package_result

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/package_group"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
)

type Result struct {
	Summary *system_result.Summary
	Kernel  []*package_group.Group
	Reboot  []*package_group.Group
	Upgrade []*package_group.Group
	Pending []*package_group.Group
	Paid    []*package_group.Group
	Open    []*package_group.Group
}
