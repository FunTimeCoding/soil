package elite

type Player struct {
	Name            string         `json:"name"`
	Account         string         `json:"account"`
	Profession      string         `json:"profession"`
	Group           int            `json:"group"`
	HasCommanderTag bool           `json:"hasCommanderTag"`
	NotInSquad      bool           `json:"notInSquad"`
	TeamIdentifier  int            `json:"teamID"`
	ActiveTimes     []int          `json:"activeTimes"`
	DPSTargets      [][]Damage     `json:"dpsTargets"`
	StatsTargets    [][]StatsEntry `json:"statsTargets"`
	Defenses        []Defense      `json:"defenses"`
	Support         []Support      `json:"support"`
	BuffUptimes     []Buff         `json:"buffUptimesActive"`
	ExtHealing      *ExtHealing    `json:"extHealingStats"`
	ExtBarrier      *ExtBarrier    `json:"extBarrierStats"`
	StatsAll        []StatsAll     `json:"statsAll"`
}
