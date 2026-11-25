package merge_request

import "gitlab.com/gitlab-org/api/client-go/v3"

func usernames(v []*gitlab.BasicUser) []string {
	var result []string

	for _, u := range v {
		if u != nil {
			result = append(result, u.Username)
		}
	}

	return result
}
