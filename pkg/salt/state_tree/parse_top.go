package state_tree

import (
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"github.com/funtimecoding/soil/pkg/salt/types/top_entry"
	"slices"
	"strings"
)

func (t *Tree) parseTop(text string) {
	environment := ""
	target := ""

	for _, line := range strings.Split(text, "\n") {
		if constant.TemplateLine.MatchString(line) {
			continue
		}

		if m := constant.EnvironmentHeader.FindStringSubmatch(line); m != nil {
			environment = m[1]

			continue
		}

		if header, isHeader := targetHeader(line); isHeader {
			target = header

			continue
		}

		item, isItem := listItem(line)

		if !isItem || environment == "" {
			continue
		}

		live := slices.Contains(t.environments, environment)

		if !live {
			t.dead[environment] = true
		}

		t.entries = append(
			t.entries,
			top_entry.New(environment, target, item, live),
		)
	}
}
