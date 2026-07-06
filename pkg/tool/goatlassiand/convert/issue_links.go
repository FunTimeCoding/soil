package convert

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func issueLinks(i *issue.Issue) []*IssueLink {
	if i.Raw.Fields.IssueLinks == nil {
		return nil
	}

	var result []*IssueLink

	for _, l := range i.Raw.Fields.IssueLinks {
		if l.OutwardIssue != nil {
			result = append(
				result,
				&IssueLink{
					Identifier: l.ID,
					Direction:  "outward",
					Type:       l.Type.Outward,
					Key:        l.OutwardIssue.Key,
					Summary:    l.OutwardIssue.Fields.Summary,
					Status:     l.OutwardIssue.Fields.Status.Name,
				},
			)
		}

		if l.InwardIssue != nil {
			result = append(
				result,
				&IssueLink{
					Identifier: l.ID,
					Direction:  "inward",
					Type:       l.Type.Inward,
					Key:        l.InwardIssue.Key,
					Summary:    l.InwardIssue.Fields.Summary,
					Status:     l.InwardIssue.Fields.Status.Name,
				},
			)
		}
	}

	return result
}
