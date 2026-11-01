package style_sheet

func New(
	offset int,
	selectors []string,
	declarations []string,
) *Rule {
	return &Rule{
		Offset:       offset,
		Selectors:    selectors,
		Declarations: declarations,
	}
}
