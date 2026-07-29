package variable_naming

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func lettersForType(t types.Type) []string {
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
		return lettersFromWord(constant.FileWord)
	case constant.PrecedenceGzipWriter:
		return []string{"z"}
	case constant.PrecedenceTarWriter:
		return []string{"t"}
	case constant.PrecedenceString:
		return lettersFromWord("string")
	case constant.PrecedenceInt:
		return []string{"i"}
	case constant.PrecedenceFloat:
		return lettersFromWord("float")
	case constant.PrecedenceBool:
		return []string{"b"}
	case constant.PrecedenceByte:
		return lettersFromWord("byte")
	case constant.PrecedenceByteSlice:
		return lettersFromWord("byte")
	case constant.PrecedenceMap:
		return lettersFromWord("map")
	case constant.PrecedenceChannel:
		return []string{"c"}
	case constant.PrecedenceStructSlice:
		return []string{"v"}
	case constant.PrecedencePrimitiveSlice:
		return lettersForPrimitiveSlice(t)
	case constant.PrecedenceStruct:
		letters := lettersFromTypeName(t)

		if letters == nil {
			return lettersFromWord("struct")
		}

		return letters
	case constant.PrecedenceInterface:
		return lettersFromTypeName(t)
	}

	return nil
}
