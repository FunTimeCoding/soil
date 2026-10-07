package metric

import "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"

func withStatus(label []string) []string {
	return WithLabel(label, constant.StatusLabel)
}
