package state_tree

import (
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"strings"
)

func (t *Tree) parseState(
	name string,
	initFile bool,
	text string,
) {
	for _, m := range constant.TemplateReference.FindAllStringSubmatch(text, -1) {
		t.addReference(name, m[1])

		if strings.HasSuffix(m[1], constant.StateExtension) {
			t.addInclude(name, stateName(m[1]))
		}
	}

	for _, m := range constant.FileReference.FindAllStringSubmatch(text, -1) {
		t.addReference(name, m[1])
		t.served[name] = append(t.served[name], m[1])
	}

	for _, m := range constant.RequisiteState.FindAllStringSubmatch(text, -1) {
		t.addInclude(name, m[1])
	}

	inside := false

	for _, line := range strings.Split(text, "\n") {
		if constant.IncludeHeader.MatchString(line) {
			inside = true

			continue
		}

		if !inside {
			continue
		}

		trimmed := strings.TrimSpace(line)

		if trimmed == "" ||
			strings.HasPrefix(trimmed, "#") ||
			constant.TemplateLine.MatchString(line) {
			continue
		}

		item, isItem := listItem(line)

		if isItem {
			t.addInclude(name, resolveInclude(name, initFile, item))

			continue
		}

		if !strings.HasPrefix(line, " ") && !strings.HasPrefix(line, "\t") {
			inside = false
		}
	}
}
