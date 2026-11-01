package style_sheet

func Parse(css string) []*Rule {
	return parseFrom(blankComments(css), 0)
}
