package palette

import "sort"

func SortResults(results []Result) {
	sort.Slice(
		results,
		func(i, j int) bool {
			if results[i].Score != results[j].Score {
				return results[i].Score > results[j].Score
			}

			return results[i].Command.Label < results[j].Command.Label
		},
	)
}
