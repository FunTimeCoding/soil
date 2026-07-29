package issue

import "github.com/funtimecoding/soil/pkg/atlassian/constant"

func ValidValue(s string) bool {
	if s == constant.JiraNilValue || s == constant.JiraUnknownValue || s == constant.JiraUnknownField {
		return false
	}

	return true
}
