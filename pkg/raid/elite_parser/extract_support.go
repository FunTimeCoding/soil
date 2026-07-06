package elite_parser

import "github.com/funtimecoding/soil/pkg/raid/elite"

func extractSupport(p *elite.Player) *Support {
	result := &Support{}

	if len(p.Support) > 0 {
		result.BoonStrips = p.Support[0].BoonStrips
		result.ConditionCleanses = p.Support[0].CondiCleanse
		result.Resurrects = p.Support[0].Resurrects
		result.ResurrectTime = p.Support[0].ResurrectTime
		result.StunBreak = p.Support[0].StunBreak
	}

	if p.ExtHealing != nil && len(p.ExtHealing.OutgoingHealing) > 0 {
		result.Healing = p.ExtHealing.OutgoingHealing[0].Healing
	}

	if p.ExtBarrier != nil && len(p.ExtBarrier.OutgoingBarrier) > 0 {
		result.Barrier = p.ExtBarrier.OutgoingBarrier[0].Barrier
	}

	return result
}
