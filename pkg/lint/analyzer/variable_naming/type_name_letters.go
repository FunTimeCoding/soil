package variable_naming

import (
	"go/types"
	"strings"
)

func typeNameLetters(t types.Type) []string {
	t = deref(t)
	named, okay := t.(*types.Named)

	if !okay {
		return nil
	}

	name := named.Obj().Name()
	lastWord := extractLastWord(name)

	return wordLetters(strings.ToLower(lastWord))
}
