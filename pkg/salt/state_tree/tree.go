package state_tree

import (
	"github.com/funtimecoding/soil/pkg/salt/types/top_entry"
	"regexp"
)

type Tree struct {
	directories  []string
	environments []string
	states       map[string]string
	files        map[string]string
	includes     map[string][]string
	statePattern map[string][]*regexp.Regexp
	references   map[string][]string
	filePattern  map[string][]*regexp.Regexp
	served       map[string][]string
	cited        []string
	citedPattern []*regexp.Regexp
	entries      []*top_entry.Entry
	dead         map[string]bool
}
