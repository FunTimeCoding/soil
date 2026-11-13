package repository

import (
	"github.com/funtimecoding/soil/pkg/errors"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func (r *Repository) siblingFile(
	sibling string,
	relative string,
) string {
	if cached, found := r.siblingContent[relative]; found {
		return cached
	}

	b, e := r.trees.Tree(r.Absolute(sibling)).Read(
		strings.TrimPrefix(relative, join.Empty(sibling, stringsConstant.Slash)),
	)
	errors.PanicOnError(e)
	result := string(b)
	r.siblingContent[relative] = result

	return result
}
