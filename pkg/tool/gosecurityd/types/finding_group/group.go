package finding_group

type Group struct {
	Vulnerability string
	Systems       []string
	NoFix         int
	PaidOnly      int
	Upgradable    int
}
