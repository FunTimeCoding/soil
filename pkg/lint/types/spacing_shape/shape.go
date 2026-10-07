package spacing_shape

type Shape struct {
	Line                string
	Number              int
	Trimmed             string
	Blank               bool
	TopLevel            bool
	ControlStart        bool
	Exit                bool
	Deferral            bool
	TopLevelDeclaration bool
	Variable            bool
	Constant            bool
	ClosingBrace        bool
	ElseContinuation    bool
	EndsWithBrace       bool
	PastTrimmed         string
	PastOpensBlock      bool
}
