package revision

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/scanner"
	"go/token"
	"strings"
)

func tokenText(source []byte) string {
	var s scanner.Scanner
	f := token.NewFileSet().AddFile("", -1, len(source))
	s.Init(f, source, nil, 0)
	var result []string

	for {
		_, t, literal := s.Scan()

		if t == token.EOF {
			break
		}

		if t == token.SEMICOLON && literal == stringsConstant.Unix {
			continue
		}

		if closesGroup(t) && len(result) > 0 &&
			result[len(result)-1] == token.COMMA.String() {
			result = result[:len(result)-1]
		}

		if literal == "" {
			literal = t.String()
		}

		result = append(result, literal)
	}

	return strings.Join(result, constant.FieldSeparator)
}
