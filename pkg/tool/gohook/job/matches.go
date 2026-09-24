package job

func (j *Job) Matches(changed []string) bool {
	if len(j.Paths) == 0 {
		return true
	}

	for _, p := range j.Paths {
		for _, c := range changed {
			if matchPath(p, c) {
				return true
			}
		}
	}

	return false
}
