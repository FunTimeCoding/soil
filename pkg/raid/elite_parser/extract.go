package elite_parser

import "github.com/funtimecoding/soil/pkg/raid/elite"

func Extract(fight *elite.Fight) []*PlayerFightStatistic {
	var result []*PlayerFightStatistic

	for _, player := range fight.Players {
		if player.NotInSquad {
			continue
		}

		stat := &PlayerFightStatistic{
			Identity:  extractIdentity(player),
			Offensive: extractOffensive(player),
			Defensive: extractDefensive(player),
			Support:   extractSupport(player),
			Boons:     extractBoons(player),
		}
		result = append(result, stat)
	}

	return result
}
