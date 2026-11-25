package service_tester

func (o *Tester) Bodies(callsign string) []string {
	o.t.Helper()
	var result []string

	for _, v := range o.connector.Notes() {
		if v.Callsign == callsign {
			result = append(result, v.Body)
		}
	}

	return result
}
