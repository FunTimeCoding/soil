package web

import "github.com/funtimecoding/soil/pkg/tool/goraidd/store"

func groupByRaid(rows []store.PlayerRaidRow) [][]store.PlayerRaidRow {
	var result [][]store.PlayerRaidRow
	var current []store.PlayerRaidRow
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
