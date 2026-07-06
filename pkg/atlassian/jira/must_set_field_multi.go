package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustSetFieldMulti(
	projectKey string,
	issueType string,
	i *jira.Issue,
	fieldName string,
	valueNames []string,
) {
	errors.PanicOnError(
		c.SetFieldMulti(
			projectKey,
			issueType,
			i,
			fieldName,
			valueNames,
		),
	)
}
