package record

type RaidPlayer struct {
	Account           string
	Name              string
	Profession        string
	Fights            int
	Damage            int
	Healing           int
	ConditionCleanses int
	BoonStrips        int
	Barrier           int
	Downs             int
	DeadCount         int
	ActiveTimeMS      int
	DistToCom         float64
}
