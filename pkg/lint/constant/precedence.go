package constant

const (
	PrecedenceError = iota
	PrecedenceTestingT
	PrecedenceTestingB
	PrecedenceReader
	PrecedenceWriter
	PrecedenceContext
	PrecedenceFile
	PrecedenceGzipWriter
	PrecedenceTarWriter
	PrecedenceString
	PrecedenceInt
	PrecedenceFloat
	PrecedenceBool
	PrecedenceByte
	PrecedenceByteSlice
	PrecedenceMap
	PrecedenceChannel
	PrecedenceStructSlice
	PrecedencePrimitiveSlice
	PrecedenceStruct
	PrecedenceInterface
	PrecedenceUnknown
)

const FileWord = "file"
