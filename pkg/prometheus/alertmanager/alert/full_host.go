package alert

import (
	"github.com/funtimecoding/soil/pkg/network"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"strings"
)

func (a *Alert) FullHost() string {
	result := a.HostDetail()
	result = strings.TrimPrefix(result, constant.SecurePrefix)

	if strings.ContainsRune(result, ':') {
		return network.SplitHost(result)
	}

	return result
}
