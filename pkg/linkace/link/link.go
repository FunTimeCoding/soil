package link

import "github.com/funtimecoding/soil/pkg/linkace/types/relation"

type Link struct {
	Identifier      int
	Title           string
	Link            string
	Description     string
	Host            string
	Status          int
	Visibility      int
	ListIdentifiers []int
	TagIdentifiers  []int
	Lists           []relation.Relation
	Tags            []relation.Relation
}
