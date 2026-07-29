package variable_naming

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func typePrecedence(t types.Type) int {
	if isErrorType(t) {
		return constant.PrecedenceError
	}

	if isNamedType(t, "testing", "T") {
		return constant.PrecedenceTestingT
	}

	if isNamedType(t, "testing", "B") {
		return constant.PrecedenceTestingB
	}

	if implementsInterface(t, "io", "Reader") {
		return constant.PrecedenceReader
	}

	if implementsInterface(t, "io", "Writer") {
		return constant.PrecedenceWriter
	}

	if isNamedType(t, "context", "Context") {
		return constant.PrecedenceContext
	}

	if isNamedType(t, "os", "File") {
		return constant.PrecedenceFile
	}

	if isNamedType(t, "compress/gzip", "Writer") {
		return constant.PrecedenceGzipWriter
	}

	if isNamedType(t, "archive/tar", "Writer") {
		return constant.PrecedenceTarWriter
	}

	underlying := deref(t).Underlying()

	if isBasicKind(underlying, types.String) {
		return constant.PrecedenceString
	}

	if isIntegerType(underlying) {
		return constant.PrecedenceInt
	}

	if isFloatType(underlying) {
		return constant.PrecedenceFloat
	}

	if isBasicKind(underlying, types.Bool) {
		return constant.PrecedenceBool
	}

	if isByteSlice(t) {
		return constant.PrecedenceByteSlice
	}

	if isBasicKind(underlying, types.Byte) {
		return constant.PrecedenceByte
	}

	if _, okay := underlying.(*types.Map); okay {
		return constant.PrecedenceMap
	}

	if _, okay := underlying.(*types.Chan); okay {
		return constant.PrecedenceChannel
	}

	if s, okay := underlying.(*types.Slice); okay {
		e := s.Elem()

		if p, okay := e.(*types.Pointer); okay {
			e = p.Elem()
		}

		if _, okay := e.Underlying().(*types.Struct); okay {
			return constant.PrecedenceStructSlice
		}

		if named, okay := e.(*types.Named); okay {
			if _, okay := named.Underlying().(*types.Struct); okay {
				return constant.PrecedenceStructSlice
			}
		}

		return constant.PrecedencePrimitiveSlice
	}

	if _, okay := underlying.(*types.Struct); okay {
		return constant.PrecedenceStruct
	}

	if _, okay := underlying.(*types.Interface); okay {
		return constant.PrecedenceInterface
	}

	return constant.PrecedenceUnknown
}
