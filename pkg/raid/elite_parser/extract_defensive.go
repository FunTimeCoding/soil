package elite_parser

import "github.com/funtimecoding/soil/pkg/raid/elite"

func extractDefensive(p *elite.Player) *Defensive {
	if len(p.Defenses) == 0 {
		return &Defensive{}
	}

	d := p.Defenses[0]

	return &Defensive{
		DamageTaken:      d.DamageTaken,
		DownCount:        d.DownCount,
		DeadCount:        d.DeadCount,
		DodgeCount:       d.DodgeCount,
		EvadedCount:      d.EvadedCount,
		BlockedCount:     d.BlockedCount,
		InvulnedCount:    d.InvulnedCount,
		InterruptedCount: d.InterruptedCount,
	}
}
