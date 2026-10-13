package go_mod

import (
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"golang.org/x/mod/modfile"
)

func SyncReplaces(
	mod []byte,
	managed []*modfile.Replace,
) ([]byte, error) {
	f, e := modfile.Parse(constant.ModFile, mod, nil)

	if e != nil {
		return nil, e
	}

	block := &modfile.LineBlock{
		Comments: modfile.Comments{
			Before: []modfile.Comment{{Token: constant.ManagedMarker}},
		},
		Token: []string{constant.ReplaceKeyword},
	}

	for _, r := range managed {
		l := &modfile.Line{Token: replaceTokens(r), InBlock: true}

		if r.Syntax != nil {
			l.Comments = r.Syntax.Comments
		}

		block.Line = append(block.Line, l)
	}

	keep := len(managed) > 0
	placed := false
	var statements []modfile.Expr

	for _, s := range f.Syntax.Stmt {
		b, okay := s.(*modfile.LineBlock)

		if !okay || !isManagedBlock(b) {
			statements = append(statements, s)

			continue
		}

		if keep && !placed {
			statements = append(statements, block)
			placed = true
		}
	}

	if keep && !placed {
		statements = append(statements, block)
	}

	f.Syntax.Stmt = statements

	return modfile.Format(f.Syntax), nil
}
