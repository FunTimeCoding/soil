package service_tester

func (t *Tester) Creates() int {
	result := 0

	for _, a := range t.Dynamic.Actions() {
		if a.GetVerb() == "create" {
			result++
		}
	}

	return result
}
