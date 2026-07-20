package elite_parser

import (
	"github.com/funtimecoding/soil/pkg/raid/constant"
	"github.com/funtimecoding/soil/pkg/raid/elite"
)

func IsValidFight(fight *elite.Fight) bool {
	if fight.DurationMS < constant.MinDuration {
		return false
	}

	allies := 0

	for _, player := range fight.Players {
		if !player.NotInSquad {
			allies++
		}
	}

	if allies < constant.MinAllies {
		return false
	}

	enemies := 0

	for _, target := range fight.Targets {
		if target.EnemyPlayer && target.TeamID != 0 {
			enemies++
		}
	}

	return enemies >= constant.MinEnemies
}
