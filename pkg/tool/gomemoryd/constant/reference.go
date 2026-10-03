package constant

const (
	ReferenceRootEnvironment = "GOMEMORY_REFERENCE_ROOT"
	MemoryScheme             = "memory://"
	LocatorSeparator         = "://"
	ProseOpening             = "([*_\"'"
	ProseClosing             = ".,;:)]*_\"'!?"
	Possessive               = "'s"
	MissingMemoryText        = "Referenced memory does not exist"
	BarePathText             = "Path in prose - wrap it in backticks so it can be checked"
	ReferenceHeader          = "References to fix:"
	CheckReferences          = "check_references"
	CheckReferencesSummary   = "Check every hand-written memory's references - paths, go: symbols, routes and memory:// citations - and list the dead ones"
	ReferencesClean          = "Every reference resolves"
)
