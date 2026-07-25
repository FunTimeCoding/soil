package convert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func JiraComments(i *issue.Issue) *[]*server.JiraComment {
	if i.Raw.Fields.Comments == nil {
		return nil
	}

	var result []*server.JiraComment

	for _, c := range i.Raw.Fields.Comments.Comments {
		comment := &server.JiraComment{
			Identifier: c.ID,
			Author:     c.Author.DisplayName,
			Body:       c.Body,
			Created:    c.Created,
		}

		if c.Updated != c.Created {
			comment.Updated = &c.Updated
		}

		result = append(result, comment)
	}

	if result == nil {
		return nil
	}

	return &result
}
