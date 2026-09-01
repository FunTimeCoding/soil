package model_context

import "github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/credential"

func describeCredentials(
	credentials []*credential.Credential,
) []map[string]any {
	var result []map[string]any

	for _, v := range credentials {
		result = append(
			result,
			map[string]any{
				"identifier":  v.Identifier,
				"path":        v.Path,
				"title":       v.Title,
				"user":        v.User,
				"locator":     v.Locator,
				"modified_at": v.ModifiedAt,
			},
		)
	}

	return result
}
