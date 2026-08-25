package raid

type PlayerFightStatistic struct {
	Identifier        uint   `gorm:"primaryKey;column:id"`
	Filename          string `gorm:"index"`
	Account           string `gorm:"index"`
	Name              string
	Profession        string
	GroupNumber       int
	Commander         bool
	ActiveTimeMS      int
	Damage            int
	DamageTaken       int
	DownCount         int
	DeadCount         int
	Kills             int
	Downs             int
	BoonStrips        int
	ConditionCleanses int
	Healing           int
	Barrier           int
	DodgeCount        int
	EvadedCount       int
	BlockedCount      int
	InvulnedCount     int
	StabilityUptime   float64
	MightUptime       float64
	FuryUptime        float64
	QuicknessUptime   float64
	ProtectionUptime  float64
	ResistanceUptime  float64
	DistToCom         float64
}
