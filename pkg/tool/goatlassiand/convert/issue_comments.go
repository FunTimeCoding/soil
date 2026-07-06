package convert

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func issueComments(i *issue.Issue) []*IssueComment {
	if i.Raw.Fields.Comments == nil {
		return nil
	}

	var result []*IssueComment

	for _, c := range i.Raw.Fields.Comments.Comments {
		comment := &IssueComment{
			Identifier: c.ID,
			Author:     c.Author.DisplayName,
			Body:       c.Body,
			Created:    c.Created,
		}

		if c.Updated != c.Created {
			comment.Updated = c.Updated
		}

		result = append(result, comment)
	}

	return result
}
