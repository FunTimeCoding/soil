package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/custom_field_value"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustCustomFieldValues(
	key string,
	issueType string,
	fieldName string,
) []custom_field_value.Value {
	result, e := c.CustomFieldValues(key, issueType, fieldName)
	errors.PanicOnError(e)

	return result
}
