package pointer

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"slices"
	"strings"
)

func Classify(
	s string,
	roots []string,
) constant.PointerClass {
	if strings.Contains(s, "://") {
		return constant.PointerClassLocator
	}

	trimmed, plugin := strings.CutPrefix(s, constant.PluginRootPrefix)

	if strings.ContainsAny(trimmed, "<>*$") {
		return constant.PointerClassPlaceholder
	}

	if strings.HasPrefix(trimmed, constant.SchemeGo) {
		return constant.PointerClassSymbol
	}

	if strings.HasPrefix(trimmed, constant.SchemeRoute) {
		return constant.PointerClassRoute
	}

	if strings.HasPrefix(trimmed, constant.SchemePath) {
		return constant.PointerClassPath
	}

	if len(trimmed) > 1 &&
		strings.HasPrefix(trimmed, constant.Quote) &&
		strings.HasSuffix(trimmed, constant.Quote) {
		return constant.PointerClassImport
	}

	if strings.HasPrefix(trimmed, constant.CommentPrefix) ||
		(strings.HasPrefix(trimmed, constant.SubstitutionPrefix) &&
			strings.Count(trimmed, "/") >= 3) {
		return constant.PointerClassPattern
	}

	if plugin {
		return constant.PointerClassRepository
	}

	if strings.HasPrefix(trimmed, "/") {
		if IsCommand(trimmed) {
			return constant.PointerClassCommand
		}

		if strings.HasPrefix(trimmed, constant.UserPathPrefix) ||
			strings.HasPrefix(trimmed, constant.HomePathPrefix) {
			return constant.PointerClassAbsolute
		}

		return constant.PointerClassSystem
	}

	if strings.ContainsAny(trimmed, constant.BraceCharacters) {
		return constant.PointerClassPlaceholder
	}

	if strings.HasPrefix(trimmed, library.ParentDirectory) {
		return constant.PointerClassSibling
	}

	root, _, _ := strings.Cut(strings.TrimPrefix(trimmed, "./"), "/")

	if !slices.Contains(roots, root) {
		return constant.PointerClassShort
	}

	return constant.PointerClassRepository
}
