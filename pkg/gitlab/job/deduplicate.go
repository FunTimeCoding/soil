package job

import "gitlab.com/gitlab-org/api/client-go/v3"

func Deduplicate(v []*gitlab.Job) []*gitlab.Job {
	var result []*gitlab.Job

	for _, e := range v {
		var found bool

		for _, existing := range result {
			if e.ID == existing.ID {
				found = true

				break
			}
		}

		if !found {
			result = append(result, e)
		}
	}

	return result
}
