package spacing

type shape struct {
	line                string
	number              int
	trimmed             string
	blank               bool
	topLevel            bool
	controlStart        bool
	exit                bool
	deferral            bool
	topLevelDeclaration bool
	variable            bool
	constant            bool
	closingBrace        bool
	elseContinuation    bool
	endsWithBrace       bool
	pastTrimmed         string
	pastOpensBlock      bool
}
