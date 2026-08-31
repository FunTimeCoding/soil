package worker

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func samePage(a *page.Page, b *page.Page) bool {
	return a.Identifier == b.Identifier &&
		a.Raw.Version.Number == b.Raw.Version.Number
}
