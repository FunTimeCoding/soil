package issue

import "github.com/funtimecoding/soil/pkg/atlassian/constant"

func Keys(s string) []string {
	return constant.JiraKeyMatch.FindAllString(s, -1)
}
