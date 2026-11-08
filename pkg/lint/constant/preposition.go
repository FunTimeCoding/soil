package constant

var (
	PrepositionAnywhere = map[string]bool{
		"of":    true,
		"by":    true,
		"for":   true,
		"at":    true,
		"until": true,
	}

	PrepositionNotFirst = map[string]bool{"from": true}

	PrepositionLast = map[string]bool{
		"to":   true,
		"into": true,
		"onto": true,
		"on":   true,
	}
)
