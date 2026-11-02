package alliance

import (
	"charm.land/lipgloss/v2/table"
	"github.com/funtimecoding/soil/pkg/gw2/log_manager/log"
	"time"
)

type allianceLog struct {
	path               string
	currentTeam        string
	matchUpStart       time.Time
	members            []string
	atRisk             []string
	aleevaFile         string
	exceptionNames     []string
	t                  *table.Table
	onTeam             []string
	onDiscord          []string
	confirmedNotOnTeam []string
	verifiedAccounts   []string
	rowCount           int
	logs               []*log.Log
	guildReport        map[string][]string
	seenDays           []*log.SeenDaysMember
	neverSeenDays      []string
}
