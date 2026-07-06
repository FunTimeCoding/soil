package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/raid"
	"github.com/funtimecoding/soil/pkg/raid/elite"
	"github.com/funtimecoding/soil/pkg/raid/elite_parser"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/time"
	"strings"
)

func (s *Store) enrichFile(
	base string,
	name string,
) {
	if !strings.HasSuffix(name, constant.DetailedWvWKillSuffix) {
		return
	}

	b := system.ReadBytes(base, name)
	var fight elite.Fight
	notation.MustDecodeBytes(b, &fight, false)
	timestamp := time.Parse(
		"2006-01-02 15:04:05 -07:00",
		fight.TimeStartStd,
	)
	alliedTeamID := 0

	if len(fight.Players) > 0 {
		alliedTeamID = fight.Players[0].TeamID
	}

	enemyCount := 0
	enemyTeams := map[int]int{}

	for _, target := range fight.Targets {
		if target.TeamID == 0 {
			continue
		}

		enemyCount++
		enemyTeams[target.TeamID]++
	}

	enemyTeamsString := string(notation.Marshal(enemyTeams))
	zevtcBase := strings.TrimSuffix(name, constant.DetailedWvWKillSuffix)
	var fightRow raid.Fight
	lookup := s.mapper.
		Where("filename LIKE ?", join.Empty("%", zevtcBase, "%")).
		First(&fightRow)

	if lookup.Error != nil {
		return
	}

	if fightRow.Enriched {
		return
	}

	s.logger.Structured("enriching_fight", "file", name, "zevtc", zevtcBase)
	s.mapper.Model(&fightRow).Updates(
		map[string]any{
			"timestamp":      timestamp,
			"duration_ms":    fight.DurationMS,
			"map_id":         fight.MapID,
			"map_name":       fight.FightName,
			"allied_count":   len(fight.Players),
			"allied_team_id": alliedTeamID,
			"enemy_count":    enemyCount,
			"enemy_teams":    enemyTeamsString,
			"success":        fight.Success,
			"enriched":       true,
		},
	)

	if !elite_parser.IsValidFight(&fight) {
		return
	}

	s.mapper.Where("filename = ?", fightRow.Filename).
		Delete(raid.NewPlayerFightStatistic())
	stats := elite_parser.Extract(&fight)

	for _, stat := range stats {
		row := raid.NewPlayerFightStatistic()
		row.Filename = fightRow.Filename
		row.Account = stat.Identity.Account
		row.Name = stat.Identity.Name
		row.Profession = stat.Identity.Profession
		row.GroupNumber = stat.Identity.Group
		row.Commander = stat.Identity.Commander
		row.ActiveTimeMS = stat.Identity.ActiveTimeMS
		row.Damage = stat.Offensive.Damage
		row.DamageTaken = stat.Defensive.DamageTaken
		row.DownCount = stat.Defensive.DownCount
		row.DeadCount = stat.Defensive.DeadCount
		row.Kills = stat.Offensive.Kills
		row.Downs = stat.Offensive.Downs
		row.BoonStrips = stat.Support.BoonStrips
		row.ConditionCleanses = stat.Support.ConditionCleanses
		row.Healing = stat.Support.Healing
		row.Barrier = stat.Support.Barrier
		row.DodgeCount = stat.Defensive.DodgeCount
		row.EvadedCount = stat.Defensive.EvadedCount
		row.BlockedCount = stat.Defensive.BlockedCount
		row.InvulnedCount = stat.Defensive.InvulnedCount
		row.StabilityUptime = stat.Boons.Stability
		row.MightUptime = stat.Boons.Might
		row.FuryUptime = stat.Boons.Fury
		row.QuicknessUptime = stat.Boons.Quickness
		row.ProtectionUptime = stat.Boons.Protection
		row.ResistanceUptime = stat.Boons.Resistance
		row.DistToCom = stat.Identity.DistToCom
		errors.PanicOnError(s.mapper.Create(row).Error)
	}

	s.logger.Structured(
		"enriched_player_stats",
		"file",
		name,
		"players",
		len(stats),
	)
}
