package elite_parser

import "github.com/funtimecoding/soil/pkg/raid/elite"

func extractIdentity(p *elite.Player) *Identity {
	result := &Identity{
		Account:    p.Account,
		Name:       p.Name,
		Profession: p.Profession,
		Group:      p.Group,
		Commander:  p.HasCommanderTag,
	}

	if len(p.ActiveTimes) > 0 {
		result.ActiveTimeMS = p.ActiveTimes[0]
	}

	if len(p.StatsAll) > 0 {
		result.DistToCom = p.StatsAll[0].DistToCom
	}

	return result
}
