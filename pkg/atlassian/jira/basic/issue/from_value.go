package issue

import "github.com/funtimecoding/soil/pkg/atlassian/constant"

func FromValue(v Values) *Issue {
	result := New()
	result.Key = v.IssueKey

	for _, e := range v.RequestFieldValues {
		switch e.FieldIdentifier {
		case constant.JiraSummaryField:
			result.Title = e.Value.(string)
		case constant.JiraDescriptionField:
			result.Body = e.Value.(string)
		}
	}

	return result
}
