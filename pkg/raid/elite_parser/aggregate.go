package elite_parser

import "github.com/funtimecoding/soil/pkg/strings/join"

func Aggregate(v []*PlayerFightStatistic) map[string]*AggregatedPlayer {
	players := map[string]*AggregatedPlayer{}

	for _, s := range v {
		key := join.Empty(
			s.Identity.Account,
			"|",
			s.Identity.Profession,
		)
		p, exists := players[key]

		if !exists {
			p = &AggregatedPlayer{
				Account:    s.Identity.Account,
				Name:       s.Identity.Name,
				Profession: s.Identity.Profession,
			}
			players[key] = p
		}

		p.Fights++
		p.TotalActiveTimeMS += s.Identity.ActiveTimeMS
		p.Offensive.Damage += s.Offensive.Damage
		p.Offensive.Kills += s.Offensive.Kills
		p.Offensive.Downs += s.Offensive.Downs
		p.Defensive.DamageTaken += s.Defensive.DamageTaken
		p.Defensive.DownCount += s.Defensive.DownCount
		p.Defensive.DeadCount += s.Defensive.DeadCount
		p.Defensive.DodgeCount += s.Defensive.DodgeCount
		p.Defensive.EvadedCount += s.Defensive.EvadedCount
		p.Defensive.BlockedCount += s.Defensive.BlockedCount
		p.Defensive.InvulnedCount += s.Defensive.InvulnedCount
		p.Defensive.InterruptedCount += s.Defensive.InterruptedCount
		p.Support.BoonStrips += s.Support.BoonStrips
		p.Support.ConditionCleanses += s.Support.ConditionCleanses
		p.Support.Healing += s.Support.Healing
		p.Support.Barrier += s.Support.Barrier
		p.Support.Resurrects += s.Support.Resurrects
		p.Support.ResurrectTime += s.Support.ResurrectTime
		p.Support.StunBreak += s.Support.StunBreak
		p.Boons.Stability += s.Boons.Stability
		p.Boons.Might += s.Boons.Might
		p.Boons.Fury += s.Boons.Fury
		p.Boons.Quickness += s.Boons.Quickness
		p.Boons.Protection += s.Boons.Protection
		p.Boons.Resistance += s.Boons.Resistance
		p.Boons.Aegis += s.Boons.Aegis
		p.Boons.Resolution += s.Boons.Resolution
		p.Boons.Swiftness += s.Boons.Swiftness
		p.Boons.Vigor += s.Boons.Vigor
		p.Boons.Regeneration += s.Boons.Regeneration
		p.Boons.Alacrity += s.Boons.Alacrity
	}

	return players
}
