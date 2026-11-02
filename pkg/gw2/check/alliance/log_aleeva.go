package alliance

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gw2"
	"github.com/funtimecoding/soil/pkg/gw2/check/alliance/aleeva_report"
	"github.com/funtimecoding/soil/pkg/strings/contains"
	"github.com/funtimecoding/soil/pkg/strings/join"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"slices"
)

func logAleeva(s *allianceLog) {
	for _, r := range aleeva_report.Parse(
		systemConstant.Temporary,
		s.aleevaFile,
	) {
		if len(r.WvwTeams) == 0 || len(r.Gw2Accounts) == 0 {
			continue
		}

		if !gw2.IsAllianceMember(s.members, r.Gw2Accounts) {
			continue
		}

		if contains.Any(r.Gw2Accounts, s.exceptionNames) {
			continue
		}

		s.verifiedAccounts = append(s.verifiedAccounts, r.Gw2Accounts...)
		var teams []string

		for _, team := range r.WvwTeams {
			teams = append(teams, team)
		}

		if slices.Contains(teams, s.currentTeam) {
			for account, teamIdentifier := range r.WvwTeams {
				if teamIdentifier == s.currentTeam {
					s.onTeam = append(s.onTeam, account)
				}
			}

			continue
		}

		s.rowCount++
		s.t.Row(r.DiscordName, join.Comma(r.Gw2Accounts), join.Comma(teams))

		for account, teamIdentifier := range r.WvwTeams {
			if teamIdentifier != s.currentTeam {
				s.confirmedNotOnTeam = append(s.confirmedNotOnTeam, account)
			}
		}

		for account, teamIdentifier := range r.WvwTeams {
			if teamIdentifier != s.currentTeam {
				s.onDiscord = append(s.onDiscord, account)
			}
		}
	}

	if s.rowCount > 0 {
		console.Format("Not on team members (%d):\n", s.rowCount)
		console.Line(s.t.Render())
	}
}
