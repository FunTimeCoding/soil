package state_tree

import "regexp"

func New(
	directories []string,
	environments []string,
) *Tree {
	result := &Tree{
		directories:  directories,
		environments: environments,
		states:       map[string]string{},
		files:        map[string]string{},
		includes:     map[string][]string{},
		statePattern: map[string][]*regexp.Regexp{},
		references:   map[string][]string{},
		filePattern:  map[string][]*regexp.Regexp{},
		served:       map[string][]string{},
		dead:         map[string]bool{},
	}
	result.load()

	return result
}
