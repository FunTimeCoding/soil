package upgrade_result

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_group"

type Result struct {
	Groups          []*upgrade_group.Group
	Hosts           int
	Vulnerabilities int
	Releases        []string
}
