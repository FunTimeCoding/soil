package resolve

import "path"

func (n *Names) ResolvePackage(importPath string) (string, error) {
	if name, found := n.names[importPath]; found {
		return name, nil
	}

	return path.Base(importPath), nil
}
