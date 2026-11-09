package alliance

import (
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gw2"
	"github.com/funtimecoding/soil/pkg/gw2/check/alliance/exceptions"
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/system/join"
	soilTime "github.com/funtimecoding/soil/pkg/time"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func Log(
	path string,
	tag string,
) {
	if tag == "" {
		tag = environment.Required(constant.AllianceEnvironment)
	}

	atRiskCutOff := environment.Required(constant.AtRiskCutOffEnvironment)
	currentTeam := environment.Required(constant.TeamEnvironment)
	start := soilTime.Parse(
		timeConstant.DateYear,
		environment.Required(constant.LinkStartDateEnvironment),
	)
	matchUpStart := time.Date(
		start.Year(),
		start.Month(),
		start.Day(),
		constant.MatchUpStartHour,
		0,
		0,
		0,
		time.UTC,
	)
	console.Format(
		"Match-up start: %s\n",
		matchUpStart.Format(timeConstant.DateMinute),
	)
	console.Format("Alliance tag: %s\n", tag)
	console.Format("Team: %s\n", currentTeam)
	console.Format("At-risk cut-off member: %s\n", atRiskCutOff)
	gw2.ImportAleevaFiles()
	members := gw2.GuildMembers(gw2.NewEnvironment(), tag)
	console.Format("Members count: %d\n", len(members))
	var atRisk []string

	for _, member := range members {
		atRisk = append(atRisk, member)

		if member == atRiskCutOff {
			break
		}
	}

	aleevaFile := gw2.LatestAleevaFile(
		system.FilesMatching(systemConstant.Temporary, constant.AleevaPrefix),
	)
	console.Format(
		"Latest Aleeva file: %s\n",
		join.Absolute(systemConstant.Temporary, aleevaFile),
	)
	var exceptionNames []string

	for _, e := range exceptions.Parse(systemConstant.Temporary) {
		exceptionNames = append(exceptionNames, e.Name)
	}

	t := table.New().Headers("Name", "Account(s)", "Team(s)").StyleFunc(
		func(
			_ int,
			_ int,
		) lipgloss.Style {
			return lipgloss.NewStyle().Padding(0, 1)
		},
	)
	s := &allianceLog{
		path:           path,
		currentTeam:    currentTeam,
		matchUpStart:   matchUpStart,
		members:        members,
		atRisk:         atRisk,
		aleevaFile:     aleevaFile,
		exceptionNames: exceptionNames,
		t:              t,
	}
	logAleeva(s)
	logSightings(s)
	logSeenDays(s)
	logLists(s)
}
