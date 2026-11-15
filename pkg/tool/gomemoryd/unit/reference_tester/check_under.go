package reference_tester

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/reference"

func CheckUnder(
	content string,
	bases ...string,
) []*reference.Finding {
	return reference.Check(content, bases, Resolver(), Known)
}
