package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
)

func initialsField(i *jira.Issue) string {
	result, _ := key_value.Dash(i.Key)

	return result
}
