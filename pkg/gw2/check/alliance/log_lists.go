package alliance

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"slices"
)

func logLists(s *allianceLog) {
	for _, e := range s.seenDays {
		labels := MemberGuilds(s.guildReport, e.Name)

		if len(labels) == 0 {
			labels = append(labels, "no-guild")
		}

		if slices.Contains(s.atRisk, e.Name) {
			labels = append(labels, "at-risk")
		}

		console.Format(
			"Seen days: %d %s %s\n",
			len(e.Days),
			e.Name,
			join.Space(labels...),
		)
	}

	for _, e := range s.neverSeenDays {
		labels := MemberGuilds(s.guildReport, e)

		if len(labels) == 0 {
			labels = append(labels, "no-guild")
		}

		if slices.Contains(s.atRisk, e) {
			labels = append(labels, "at-risk")
		}

		console.Format("Seen days: 0 %s %s\n", e, join.Space(labels...))
	}
}
