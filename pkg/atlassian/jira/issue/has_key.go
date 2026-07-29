package issue

import "github.com/funtimecoding/soil/pkg/atlassian/constant"

func HasKey(s string) bool {
	return constant.JiraKeyMatch.MatchString(s)
}
