package outage_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func verdict(
	name string,
	want constant.OutageReach,
	failed bool,
	outage bool,
	reached bool,
	captured int,
	answer string,
) []string {
	var result []string

	if captured > 0 {
		result = append(result, fmt.Sprintf("%s: captured %d", name, captured))
	}

	switch want {
	case constant.OutageReaches:
		if !reached {
			result = append(
				result,
				fmt.Sprintf("%s: never reached the peer", name),
			)
		} else if !failed {
			result = append(
				result,
				fmt.Sprintf("%s: answered through the outage", name),
			)
		} else if !outage {
			result = append(
				result,
				fmt.Sprintf("%s: failed with another answer: %s", name, answer),
			)
		}
	case constant.OutageDegrades:
		if !reached {
			result = append(
				result,
				fmt.Sprintf("%s: never reached the peer", name),
			)
		} else if failed {
			result = append(
				result,
				fmt.Sprintf("%s: failed instead of degrading", name),
			)
		}
	case constant.OutageElsewhere:
		if reached {
			result = append(result, fmt.Sprintf("%s: reached the peer", name))
		}
	}

	return result
}
