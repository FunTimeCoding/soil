package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/embed_model"
)

func embedModelByName(name string) *embed_model.Model {
	var known []string

	for _, m := range constant.EmbedModels {
		if m.Name == name {
			return m
		}

		known = append(known, m.Name)
	}

	panic(
		fmt.Sprintf(
			"unknown embed model %q, known: %s",
			name,
			join.CommaSpace(known),
		),
	)
}
