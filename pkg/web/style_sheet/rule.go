package style_sheet

type Rule struct {
	Offset       int
	Selectors    []string
	Declarations []string
}
