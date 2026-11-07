package permission

import "sort"

func Compute(
	prompts []*Prompt,
	sessions int,
	days int,
) *Report {
	result := &Report{Days: days, Sessions: sessions}
	rows := map[string]*Row{}
	reached := map[string]map[string]bool{}
	reasons := map[string]map[string]int{}
	samples := map[string][]*Prompt{}

	for _, p := range prompts {
		result.Prompts++

		if p.Declined {
			result.Declined++
		}

		if p.Cleared {
			result.Cleared++
		}

		for _, b := range p.Blockers {
			if rows[b] == nil {
				rows[b] = &Row{Blocker: b}
				reached[b] = map[string]bool{}
				reasons[b] = map[string]int{}
			}

			r := rows[b]
			r.Prompts++
			reached[b][p.Session] = true
			reasons[b][p.Reason]++
			samples[b] = append(samples[b], p)

			if p.When.After(r.Last) {
				r.Last = p.When
			}
		}
	}

	for b, r := range rows {
		r.Sessions = len(reached[b])
		r.Reason = mostCommon(reasons[b])
		r.Commands = newestCommands(samples[b])
		result.Rows = append(result.Rows, r)
	}

	sort.Slice(
		result.Rows,
		func(i, j int) bool {
			if result.Rows[i].Prompts != result.Rows[j].Prompts {
				return result.Rows[i].Prompts > result.Rows[j].Prompts
			}

			return result.Rows[i].Blocker < result.Rows[j].Blocker
		},
	)

	return result
}
