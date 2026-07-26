package elite_parser

import (
	"github.com/funtimecoding/soil/pkg/raid/constant"
	"github.com/funtimecoding/soil/pkg/raid/elite"
)

func extractBoons(p *elite.Player) *Boons {
	result := &Boons{}

	for _, b := range p.BuffUptimes {
		if len(b.Entries) == 0 {
			continue
		}

		uptime := b.Entries[0].Uptime

		switch b.Identifier {
		case constant.BuffStability:
			result.Stability = uptime
		case constant.BuffMight:
			result.Might = uptime
		case constant.BuffFury:
			result.Fury = uptime
		case constant.BuffQuickness:
			result.Quickness = uptime
		case constant.BuffProtection:
			result.Protection = uptime
		case constant.BuffResistance:
			result.Resistance = uptime
		case constant.BuffAegis:
			result.Aegis = uptime
		case constant.BuffResolution:
			result.Resolution = uptime
		case constant.BuffSwiftness:
			result.Swiftness = uptime
		case constant.BuffVigor:
			result.Vigor = uptime
		case constant.BuffRegeneration:
			result.Regeneration = uptime
		case constant.BuffAlacrity:
			result.Alacrity = uptime
		}
	}

	return result
}
