package host_mapping

import "github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"

func configurationAddress(address string) string {
	return constant.InstanceKey.ReplaceAllString(address, "")
}
