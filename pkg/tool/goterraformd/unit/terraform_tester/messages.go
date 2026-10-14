package terraform_tester

func (o *Tester) Messages() []string {
	var result []string

	for _, v := range o.Reporter.Events() {
		result = append(result, v.Error.Error())
	}

	return result
}
