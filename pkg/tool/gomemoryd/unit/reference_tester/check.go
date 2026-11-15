package reference_tester

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/reference"

func Check(content string) []*reference.Finding {
	return CheckUnder(content)
}
