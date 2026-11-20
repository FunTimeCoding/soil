package service_tester

func (o *Tester) Messages() []string {
	o.t.Helper()
	o.Service()
	var result []string

	for _, v := range o.Reporter.Events() {
		result = append(result, v.Error.Error())
	}

	return result
}
