package elite_parser

import "github.com/funtimecoding/soil/pkg/raid/elite"

func extractOffensive(p *elite.Player) *Offensive {
	result := &Offensive{}

	for _, targets := range p.DPSTargets {
		for _, t := range targets {
			result.Damage += t.Damage
		}
	}

	for _, t := range p.StatsTargets {
		for _, t := range t {
			result.Kills += t.Killed
			result.Downs += t.Downed
		}
	}

	return result
}
