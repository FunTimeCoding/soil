package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
)

func snippetPath(name string) (string, error) {
	if name == "" {
		return "", validation.New("name is required")
	}

	if !bareName(name) {
		return "", validation.New("not a bare name: %s", name)
	}

	return join.Slash([]string{constant.SnippetDirectory, name}), nil
}
