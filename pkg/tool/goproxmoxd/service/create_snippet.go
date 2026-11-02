package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (s *Service) CreateSnippet(
	instance string,
	name string,
	content string,
) (string, error) {
	path, e := snippetPath(name)

	if e != nil {
		return "", e
	}

	if content == "" {
		return "", validation.New("content is required")
	}

	c, e := s.shells.Client(instance)

	if e != nil {
		return "", e
	}

	if f := c.WriteFile(path, []byte(content)); f != nil {
		return "", s.forgetDropped(instance, c, f)
	}

	return fmt.Sprintf("local:snippets/%s", name), nil
}
