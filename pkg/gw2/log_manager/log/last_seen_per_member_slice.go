package log

import (
	"sort"
	"time"
)

func LastSeenPerMemberSlice(
	members []string,
	logs []*Log,
	since *time.Time,
) []LastSeenMember {
	var result []LastSeenMember

	for _, member := range members {
		if t := lastSeen(logs, member); t != nil {
			if since != nil && t.Before(*since) {
				continue
			}

			result = append(result, LastSeenMember{Name: member, Time: *t})
		}
	}

	sort.SliceStable(
		result,
		func(
			i int,
			j int,
		) bool {
			return result[i].Time.Before(result[j].Time)
		},
	)

	return result
}
