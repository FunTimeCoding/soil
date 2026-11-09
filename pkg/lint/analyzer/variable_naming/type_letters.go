package variable_naming

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func typeLetters(t types.Type) []string {
	p := typePrecedence(t)

	switch p {
	case constant.PrecedenceError:
		return []string{"e", "f", "g", "h"}
	case constant.PrecedenceTestingT:
		return []string{"t"}
	case constant.PrecedenceTestingB:
		return []string{"b"}
	case constant.PrecedenceReader:
		return []string{"r"}
	case constant.PrecedenceWriter:
		return []string{"w"}
	case constant.PrecedenceContext:
		return []string{"x"}
	case constant.PrecedenceFile:
		return wordLetters(constant.FileWord)
	case constant.PrecedenceGzipWriter:
		return []string{"z"}
	case constant.PrecedenceTarWriter:
		return []string{"t"}
	case constant.PrecedenceString:
		return wordLetters("string")
	case constant.PrecedenceInt:
		return []string{"i"}
	case constant.PrecedenceFloat:
		return wordLetters("float")
	case constant.PrecedenceBool:
		return []string{"b"}
	case constant.PrecedenceByte:
		return wordLetters("byte")
	case constant.PrecedenceByteSlice:
		return wordLetters("byte")
	case constant.PrecedenceMap:
		return wordLetters("map")
	case constant.PrecedenceChannel:
		return []string{"c"}
	case constant.PrecedenceStructSlice:
		return []string{"v"}
	case constant.PrecedencePrimitiveSlice:
		return primitiveSliceLetters(t)
	case constant.PrecedenceStruct:
		letters := typeNameLetters(t)

		if letters == nil {
			return wordLetters("struct")
		}

		return letters
	case constant.PrecedenceInterface:
		return typeNameLetters(t)
	}

	return nil
}
