package jira

import "github.com/funtimecoding/soil/pkg/atlassian/jira/field_map"

func (c *Client) FieldMap() (*field_map.Map, error) {
	if c.fieldMap != nil {
		return c.fieldMap, nil
	}

	fields, e := c.Fields()

	if e != nil {
		return nil, e
	}

	c.fieldMap = field_map.New(fields)

	return c.fieldMap, nil
}
