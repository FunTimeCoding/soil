package store

import "time"

type PlayerRaidRow struct {
	RaidIdentifier    uint `gorm:"column:raid_id"`
	RaidName          string
	RaidDate          time.Time
	RaidFights        int
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
