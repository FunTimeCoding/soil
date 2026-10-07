package web

import "github.com/funtimecoding/soil/pkg/tool/goraidd/store/record"

func groupByRaid(rows []record.PlayerRaid) [][]record.PlayerRaid {
	var result [][]record.PlayerRaid
	var current []record.PlayerRaid
	var currentIdentifier uint

	for _, r := range rows {
		if r.RaidIdentifier != currentIdentifier {
			if len(current) > 0 {
				result = append(result, current)
			}

			current = nil
			currentIdentifier = r.RaidIdentifier
		}

		current = append(current, r)
	}

	if len(current) > 0 {
		result = append(result, current)
	}

	return result
}
