package repository

func (r *Repository) testDirectories(name string) map[string]bool {
	if r.tests == nil {
		r.tests = r.testNames()
	}

	return r.tests[name]
}
