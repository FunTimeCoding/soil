package section

type Section struct {
	Level     int
	Title     string
	FirstLine int
	LastLine  int
	Text      string
	Blocks    []*Block
}
