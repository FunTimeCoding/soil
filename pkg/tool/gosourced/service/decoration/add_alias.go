package decoration

import "github.com/dave/dst"

func (s *Set) AddAlias(
	file *dst.File,
	importPath string,
	alias string,
) {
	if s.Aliases[file] == nil {
		s.Aliases[file] = make(map[string]string)
	}

	s.Aliases[file][importPath] = alias
}
