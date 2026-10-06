package comparison

func Compare(
	before string,
	after string,
) *Comparison {
	from, fromOrder := wordCounts(before)
	to, toOrder := wordCounts(after)
	seen := map[string]bool{}
	result := &Comparison{}

	for _, title := range append(fromOrder, toOrder...) {
		if seen[title] {
			continue
		}

		seen[title] = true
		c := &Change{Title: title}
		c.Removed, c.RemovedCount = wordDifference(from[title], to[title])
		c.Added, c.AddedCount = wordDifference(to[title], from[title])

		if c.RemovedCount > 0 || c.AddedCount > 0 {
			result.Changes = append(result.Changes, c)
		}
	}

	result.Removed, result.Added = overallDifference(from, to)

	return result
}
