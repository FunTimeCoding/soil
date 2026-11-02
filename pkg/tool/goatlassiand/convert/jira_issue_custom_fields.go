package convert

import (
	"github.com/andygrunwald/go-jira"
	atlassianConstant "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func JiraIssueCustomFields(
	i *issue.Issue,
	t *jira.MetaIssueType,
	includeComments bool,
) *IssueWithCustomFields {
	result := &IssueWithCustomFields{JiraIssue: JiraIssue(i)}
	result.Links = issueLinks(i)

	if includeComments {
		result.Comments = issueComments(i)
	}

	if t != nil {
		fields := make(map[string]string)

		for key := range t.Fields {
			if !strings.HasPrefix(key, "customfield_") {
				continue
			}

			name, e := t.Fields.String(join.Empty(key, "/name"))

			if e != nil {
				continue
			}

			value := i.CustomValue(name)

			if value == "" ||
				value == atlassianConstant.JiraNilValue ||
				value == atlassianConstant.JiraUnknownField ||
				value == atlassianConstant.JiraUnknownValue {
				continue
			}

			fields[name] = value
		}

		if len(fields) > 0 {
			result.CustomFields = fields
		}
	}

	return result
}
