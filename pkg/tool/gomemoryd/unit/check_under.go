package unit

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/reference"

func checkUnder(
	content string,
	bases ...string,
) []*reference.Finding {
	return reference.Check(content, bases, stubResolver(), knownMemory)
}
