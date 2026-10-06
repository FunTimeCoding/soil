package service

import "golang.org/x/tools/go/packages"

func checkCollision(
	p *packages.Package,
	targetName string,
	receiver string,
) error {
	if receiver == "" {
		return checkScopeCollision(p, targetName)
	}

	return checkMemberCollision(p, targetName, receiver)
}
