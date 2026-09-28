package label

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"strings"
)

func ServiceNames(v *[]*client.Label) []string {
	var result []string

	if v == nil {
		return result
	}

	for _, l := range *v {
		if !strings.HasPrefix(l.Key, constant.ServiceLabelPrefix) {
			continue
		}

		if name := strings.TrimPrefix(
			l.Key,
			constant.ServiceLabelPrefix,
		); name != "" {
			result = append(result, name)
		}
	}

	return result
}
