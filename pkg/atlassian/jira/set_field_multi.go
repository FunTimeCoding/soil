package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/custom_field_value"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"slices"
)

func (c *Client) SetFieldMulti(
	projectKey string,
	issueType string,
	i *jira.Issue,
	fieldName string,
	valueNames []string,
) error {
	valid, e := c.CustomFieldValues(projectKey, issueType, fieldName)

	if e != nil {
		return e
	}

	if len(valid) == 0 {
		return fmt.Errorf(
			"project %s type %s field %s has no values",
			projectKey,
			issueType,
			fieldName,
		)
	}

	var validNames []string

	for _, u := range valid {
		validNames = append(validNames, u.Value)
	}

	var notFound []string

	for _, n := range valueNames {
		if !slices.Contains(validNames, n) {
			notFound = append(notFound, n)
		}
	}

	if len(notFound) > 0 {
		return fmt.Errorf(
			"field %s value(s) not found: %s, valid: %s",
			fieldName,
			join.Comma(notFound),
			join.Comma(validNames),
		)
	}

	var result []custom_field_value.Value

	for _, u := range valid {
		if slices.Contains(valueNames, u.Value) {
			result = append(result, u)
		}
	}

	return c.SetField(i, fieldName, result)
}
