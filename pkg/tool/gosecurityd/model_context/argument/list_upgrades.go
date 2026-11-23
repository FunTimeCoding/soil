package argument

type ListUpgrades struct {
	Selection
	Release string  `json:"release"`
	Kernel  bool    `json:"kernel"`
	Score   float64 `json:"score"`
}
