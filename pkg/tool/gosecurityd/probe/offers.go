package probe

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"strconv"
	"strings"
)

func offers(lines []string) []client.KernelOffer {
	var result []client.KernelOffer

	for _, line := range lines {
		if m := constant.OfferPattern.FindStringSubmatch(line); m != nil {
			priority, e := strconv.Atoi(m[2])

			if e != nil {
				continue
			}

			result = append(
				result,
				client.KernelOffer{
					Version:  m[1],
					Priority: priority,
					Origins:  []string{},
				},
			)

			continue
		}

		if m := constant.OriginPattern.FindStringSubmatch(line); m != nil &&
			len(result) > 0 {
			origin := strings.Fields(m[2])

			if len(origin) > 2 && origin[len(origin)-1] == constant.IndexSuffix {
				origin = origin[:len(origin)-2]
			}

			last := &result[len(result)-1]
			last.Origins = append(last.Origins, join.Space(origin...))
		}
	}

	return result
}
