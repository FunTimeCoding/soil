package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

func add(
	seen map[int64]*door,
	loaded map[int64]bool,
	identifier int64,
	name string,
	edge *client.Relation,
	otherIdentifier int64,
	otherName string,
) {
	if loaded[identifier] || !loaded[otherIdentifier] {
		return
	}

	if _, found := seen[identifier]; found {
		return
	}

	kind := "untyped"

	if edge.Type != nil && *edge.Type != "" {
		kind = *edge.Type
	}

	seen[identifier] = &door{
		Identifier: identifier,
		Name:       name,
		Relation:   kind,
		Source:     otherName,
	}
}
