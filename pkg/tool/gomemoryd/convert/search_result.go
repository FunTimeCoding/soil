package convert

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func SearchResult(r *store.SearchResult) *SlimSearchResult {
	return &SlimSearchResult{
		Identifier:       r.Identifier,
		Name:             r.Name,
		Content:          r.Content,
		Description:      r.Description,
		Scope:            r.Scope,
		Tags:             r.Tags,
		Rank:             r.Rank,
		ParentIdentifier: r.ParentIdentifier,
		ParentName:       r.ParentName,
	}
}
