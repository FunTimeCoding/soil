package permission

type Report struct {
	Days     int
	Sessions int
	Prompts  int
	Declined int
	Cleared  int
	Rows     []*Row
}
