package constant

var (
	PrepositionAnywhere = map[string]bool{
		"of":      true,
		"by":      true,
		"for":     true,
		"at":      true,
		"until":   true,
		"in":      true,
		"through": true,
		"after":   true,
		"since":   true,
		"before":  true,
		"without": true,
		"between": true,
		"among":   true,
		"toward":  true,
		"inside":  true,
		"outside": true,
	}

	PrepositionNotFirst = map[string]bool{"from": true, "with": true}

	PrepositionLast = map[string]bool{
		"to":   true,
		"into": true,
		"onto": true,
		"on":   true,
	}

	MomentSuffix = map[string]bool{
		"column": true,
		"field":  true,
		"index":  true,
	}
)

const QuestionSegment = "in"
