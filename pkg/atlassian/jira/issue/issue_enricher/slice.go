package issue_enricher

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

type Slice func(*issue.Issue) []string
