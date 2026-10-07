package link

import (
	"github.com/funtimecoding/soil/pkg/linkace/response"
	"github.com/funtimecoding/soil/pkg/linkace/types/relation"
)

func New(
	v response.Link,
	host string,
) *Link {
	lists := make([]relation.Relation, len(v.Lists))
	listIdentifiers := make([]int, len(v.Lists))

	for i, l := range v.Lists {
		lists[i] = relation.Relation{Identifier: l.Identifier, Name: l.Name}
		listIdentifiers[i] = l.Identifier
	}

	tags := make([]relation.Relation, len(v.Tags))
	tagIdentifiers := make([]int, len(v.Tags))

	for i, t := range v.Tags {
		tags[i] = relation.Relation{Identifier: t.Identifier, Name: t.Name}
		tagIdentifiers[i] = t.Identifier
	}

	return &Link{
		Identifier:      v.Identifier,
		Title:           v.Title,
		Link:            v.Link,
		Description:     v.Description,
		Host:            host,
		Status:          v.Status,
		Visibility:      v.Visibility,
		ListIdentifiers: listIdentifiers,
		TagIdentifiers:  tagIdentifiers,
		Lists:           lists,
		Tags:            tags,
	}
}
