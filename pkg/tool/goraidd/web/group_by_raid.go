package web

import "github.com/funtimecoding/soil/pkg/tool/goraidd/store"

func groupByRaid(rows []store.PlayerRaidRow) [][]store.PlayerRaidRow {
	var result [][]store.PlayerRaidRow
	var current []store.PlayerRaidRow
	var currentID uint

	for _, r := range rows {
		if r.RaidID != currentID {
			if len(current) > 0 {
				result = append(result, current)
			}

			current = nil
			currentID = r.RaidID
		}

		current = append(current, r)
	}

	if len(current) > 0 {
		result = append(result, current)
	}

	return result
}
