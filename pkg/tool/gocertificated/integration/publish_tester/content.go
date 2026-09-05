package publish_tester

import "gitlab.com/gitlab-org/api/client-go/v2"

func Content(
	v []*gitlab.CommitActionOptions,
	path string,
) string {
	for _, a := range v {
		if *a.FilePath == path {
			return *a.Content
		}
	}

	return ""
}
