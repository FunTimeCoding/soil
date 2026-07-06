package elite_parser

import "github.com/funtimecoding/soil/pkg/raid/elite"

func extractBoons(p *elite.Player) *Boons {
	result := &Boons{}

	for _, b := range p.BuffUptimes {
		if len(b.Entries) == 0 {
			continue
		}

		uptime := b.Entries[0].Uptime

		switch b.Identifier {
		case elite.BuffStability:
			result.Stability = uptime
		case elite.BuffMight:
			result.Might = uptime
		case elite.BuffFury:
			result.Fury = uptime
		case elite.BuffQuickness:
			result.Quickness = uptime
		case elite.BuffProtection:
			result.Protection = uptime
		case elite.BuffResistance:
			result.Resistance = uptime
		case elite.BuffAegis:
			result.Aegis = uptime
		case elite.BuffResolution:
			result.Resolution = uptime
		case elite.BuffSwiftness:
			result.Swiftness = uptime
		case elite.BuffVigor:
			result.Vigor = uptime
		case elite.BuffRegeneration:
			result.Regeneration = uptime
		case elite.BuffAlacrity:
			result.Alacrity = uptime
		}
	}

	return result
}
