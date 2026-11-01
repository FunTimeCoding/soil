package constant

const (
	MemoryHeadingPrefix = "## "
	IdentifierOpen      = " ("
	IdentifierClose     = ")"
	TagOpen             = " ["
	TagClose            = "]"
	EntrySeparator      = " - "
	ChildIndent         = "    + "
	RankSeparator       = " rank "
	BasePrefix          = "Base: "
	Backtick            = "`"
)

const (
	SectionPrefix            = "# "
	AlwaysSectionHeading     = "# always"
	IndexSectionHeading      = "# index"
	RelevantSectionHeading   = "# relevant"
	ImpressionSectionHeading = "# impressions"
	CompletionSectionHeading = "# completions"
)

const (
	QueryAwayNotice   = "goqueryd is away - topic-matched memories and completions left out"
	QueryFailedNotice = "goqueryd failed - topic-matched memories and completions left out"
)
