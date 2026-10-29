package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/link_type"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/link_type_response"
)

func (c *Client) LinkTypes() ([]link_type.Type, error) {
	var parsed link_type_response.Response

	if e := c.basic.GetPath(constant.JiraLinkTypes, &parsed); e != nil {
		return nil, e
	}

	return parsed.IssueLinkTypes, nil
}
