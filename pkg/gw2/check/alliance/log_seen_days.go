package alliance

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gw2/log_manager/log"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"slices"
	"time"
)

func logSeenDays(s *allianceLog) {
	daysAgo := 56
	ago := time.Now().AddDate(0, 0, -daysAgo)
	s.seenDays = log.SeenDaysPerMemberSlice(s.members, s.logs, &ago)
	console.Format("Total members: %d\n", len(s.members))
	console.Format(
		"  Seen days count since %s (%d days): %d\n",
		ago.Format(timeConstant.DateMinute),
		daysAgo,
		len(s.seenDays),
	)
	var seenMembers []string
	var lessThanFour []string

	for _, e := range s.seenDays {
		seenMembers = append(seenMembers, e.Name)

		if len(e.Days) < 4 {
			lessThanFour = append(lessThanFour, e.Name)
		}
	}

	console.Format("  Less than four days count: %d\n", len(lessThanFour))
	var exceptionCount int
	var noAleevaOrNotOnTeamCount int
	var confirmedNotOnTeamCount int

	for _, member := range s.members {
		if !slices.Contains(s.onTeam, member) {
			noAleevaOrNotOnTeamCount++
		}

		if slices.Contains(s.confirmedNotOnTeam, member) {
			confirmedNotOnTeamCount++
		}

		if slices.Contains(s.exceptionNames, member) {
			exceptionCount++

			continue
		}

		if !slices.Contains(seenMembers, member) {
			s.neverSeenDays = append(s.neverSeenDays, member)
		}
	}

	console.Format("  Never seen days count: %d\n", len(s.neverSeenDays))
	console.Format("    Exception count: %d\n", exceptionCount)
	console.Format(
		"    Confirmed not on team count: %d\n",
		confirmedNotOnTeamCount,
	)
	console.Format(
		"    No Aleeva count: %d\n",
		noAleevaOrNotOnTeamCount-confirmedNotOnTeamCount,
	)
}
