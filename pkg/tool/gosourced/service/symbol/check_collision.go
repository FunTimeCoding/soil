package symbol

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/identifier"
	"golang.org/x/tools/go/packages"
)

func checkCollision(
	p *packages.Package,
	targetName string,
	receiver string,
) error {
	if receiver == "" {
		return identifier.CheckScopeCollision(p, targetName)
	}

	return checkMemberCollision(p, targetName, receiver)
}
