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
	BaseParameterDescription = "Directory the memory's bare file names resolve under, comma-separated for several (e.g. '../soil/doc/ai/spec/error-handling'); `mcp.md` is then checked as that directory's file, and the base is shown as the body's first line"
	RewrittenPrefix          = "Rewrote citations in: "
	UpdateMemoryDescription  = "Update an existing memory. Records the previous version in history. An omitted or empty name, content, description or base keeps its stored value, and tags are always kept, so a description can change alone; clear_base removes the base. A new name rewrites every memory:// citation to the old one, and is refused if the scope already has it."
	BaseUpdateDescription    = "Directory the memory's bare file names resolve under, comma-separated for several; omit or leave empty to keep the current base"
	ClearBaseDescription     = "Remove the memory's base; wins over base when both are given"
	ReferencesClean          = "Every reference resolves"
)
