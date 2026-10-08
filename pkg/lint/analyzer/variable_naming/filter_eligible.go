package variable_naming

func filterEligible(variables []TypedVariable) []TypedVariable {
	var result []TypedVariable

	for _, v := range variables {
		if isEligible(v) {
			result = append(result, v)
		}
	}

	return result
}
