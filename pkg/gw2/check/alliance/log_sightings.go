package alliance

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gw2"
	"github.com/funtimecoding/soil/pkg/gw2/check/alliance/guilds"
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/gw2/log_manager/log"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"slices"
	"sort"
)

func logSightings(s *allianceLog) {
	console.Format("Members: %s\n", join.Comma(s.members))
	s.logs = log.NewSlice(
		gw2.ParseLogs(system.ReadBytes(s.path, constant.LogFile), false),
	)
	var unverified []string

	for _, member := range s.members {
		if !slices.Contains(s.verifiedAccounts, member) {
			unverified = append(unverified, member)
		}
	}

	seen := log.LastSeenPerMemberSlice(s.members, s.logs, &s.matchUpStart)
	console.Format("Last seen since relink count: %d\n", len(seen))

	for _, v := range seen {
		console.Format(
			"Last seen: %s: %s\n",
			v.Name,
			v.Time.Format(timeConstant.DateMinute),
		)
	}

	var neverSeen []string
	var foundExceptions []string
	seenMap := log.LastSeenPerMemberMap(s.members, s.logs, &s.matchUpStart)

	for _, member := range s.members {
		if _, okay := seenMap[member]; !okay {
			if slices.Contains(s.exceptionNames, member) {
				foundExceptions = append(foundExceptions, member)
			} else {
				neverSeen = append(neverSeen, member)
			}
		} else {
			if slices.Contains(s.exceptionNames, member) {
				console.Format("Found anyway exception: %s\n", member)
			}
		}
	}

	var uselessException []string

	for _, e := range s.exceptionNames {
		if !slices.Contains(s.members, e) {
			uselessException = append(uselessException, e)
			console.Format("Useless exception: %s\n", e)
		}
	}

	console.Format("Useless exceptions: %d\n", len(uselessException))
	s.guildReport = guilds.Parse(systemConstant.Temporary)
	sort.Strings(foundExceptions)
	console.Format(
		"Exceptions (%d): %s\n",
		len(foundExceptions),
		join.Comma(foundExceptions),
	)
	console.Format("Never seen count: %d\n", len(neverSeen))
	sort.Strings(neverSeen)

	for _, name := range neverSeen {
		labels := GuildsOfMember(s.guildReport, name)

		if len(labels) == 0 {
			labels = append(labels, "no-guild")
		}

		if slices.Contains(s.onTeam, name) {
			labels = append(labels, "on-current-team")
		}

		if slices.Contains(s.onDiscord, name) {
			labels = append(labels, "on-discord")
		}

		if slices.Contains(unverified, name) {
			labels = append(labels, "unverified")
		}

		if slices.Contains(s.atRisk, name) {
			labels = append(labels, "at-risk")
		}

		console.Format("Never seen: %s %s\n", name, join.Space(labels...))
	}
}
