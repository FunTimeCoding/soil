package collector

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"

func withStatus(label []string) []string {
	return withLabel(label, constant.StatusLabel)
}
