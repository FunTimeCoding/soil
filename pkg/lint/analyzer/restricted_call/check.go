package restricted_call

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"golang.org/x/tools/go/packages"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	CheckRules(p, results, rules)
}
