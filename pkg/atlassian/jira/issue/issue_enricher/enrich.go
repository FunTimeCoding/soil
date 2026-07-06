package issue_enricher

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (e *Enricher) enrich(i *issue.Issue) {
	if e.concerns != nil {
		i.SetConcerns(e.concerns(i))
	}

	if e.score != nil {
		i.SetScore(e.score(i))
	}

	if len(e.commentNameFilter) > 0 {
		i.SetCommentNameFilter(e.commentNameFilter)
	}
}
