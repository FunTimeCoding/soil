package publish_tester

import "gitlab.com/gitlab-org/api/client-go/v2"

func Paths(v []*gitlab.CommitActionOptions) []string {
	result := make([]string, 0, len(v))

	for _, a := range v {
		result = append(result, *a.FilePath)
	}

	return result
}
