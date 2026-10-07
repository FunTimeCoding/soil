package pointer

import (
	"github.com/funtimecoding/soil/pkg/constant"
	constant1 "github.com/funtimecoding/soil/pkg/lint/constant"
	"slices"
	"strings"
)

func Classify(
	s string,
	roots []string,
) constant1.PointerClass {
	trimmed, plugin := strings.CutPrefix(s, constant1.PluginRootPrefix)

	if strings.ContainsAny(trimmed, "<>*$") {
		return constant1.PointerClassPlaceholder
	}

	if strings.Contains(s, constant1.LocatorSeparator) {
		return constant1.PointerClassLocator
	}

	if strings.HasPrefix(trimmed, constant1.SchemeGo) {
		return constant1.PointerClassSymbol
	}

	if strings.HasPrefix(trimmed, constant1.SchemeRoute) {
		return constant1.PointerClassRoute
	}

	if strings.HasPrefix(trimmed, constant1.SchemePath) {
		return constant1.PointerClassPath
	}

	if len(trimmed) > 1 &&
		strings.HasPrefix(trimmed, constant1.Quote) &&
		strings.HasSuffix(trimmed, constant1.Quote) {
		return constant1.PointerClassImport
	}

	if strings.HasPrefix(trimmed, constant1.CommentPrefix) ||
		(strings.HasPrefix(trimmed, constant1.SubstitutionPrefix) &&
			strings.Count(trimmed, "/") >= 3) {
		return constant1.PointerClassPattern
	}

	if plugin {
		return constant1.PointerClassRepository
	}

	if strings.HasPrefix(trimmed, "/") {
		if IsCommand(trimmed) {
			return constant1.PointerClassCommand
		}

		if strings.HasPrefix(trimmed, constant1.UserPathPrefix) ||
			strings.HasPrefix(trimmed, constant1.HomePathPrefix) {
			return constant1.PointerClassAbsolute
		}

		return constant1.PointerClassSystem
	}

	if strings.ContainsAny(trimmed, constant1.BraceCharacters) {
		return constant1.PointerClassPlaceholder
	}

	if strings.HasPrefix(trimmed, constant.ParentDirectory) {
		return constant1.PointerClassSibling
	}

	root, _, _ := strings.Cut(strings.TrimPrefix(trimmed, "./"), "/")

	if !slices.Contains(roots, root) {
		return constant1.PointerClassShort
	}

	return constant1.PointerClassRepository
}
