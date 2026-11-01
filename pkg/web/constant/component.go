package constant

const (
	SectionClass           = "section"
	ActionClass            = "action"
	ActionSmallClass       = "small"
	NumericClass           = "numeric"
	AbsentClass            = "absent"
	ClusterClass           = "cluster"
	FactsClass             = "facts"
	RowScope               = "row"
	TileClass              = "tile"
	TileGridClass          = "tile-grid"
	DialClass              = "dial"
	DialCenterClass        = "dial-center"
	DialNeedleClass        = "dial-needle"
	DialPointClass         = "dial-point"
	DialTickClass          = "dial-tick"
	DialAngleVariable      = "--dial-angle"
	DialBackgroundVariable = "--dial-background"
	StatusDotClass         = "status-dot"
	PagerClass             = "pager"
	PagerPositionClass     = "pager-position"
	SearchAttribute        = "data-search"
	BodyTarget             = "body"
)

const (
	DialDegrees = 360.0
	DialQuarter = 90.0
)

const (
	OkayColor     = "var(--okay)"
	WarnColor     = "var(--warn)"
	BadColor      = "var(--bad)"
	ConflictColor = "var(--conflict)"
	QuietColor    = "var(--quiet)"
	BusyColor     = "var(--busy)"
)

const (
	OkayVariable     = "--okay"
	WarnVariable     = "--warn"
	BadVariable      = "--bad"
	ConflictVariable = "--conflict"
	QuietVariable    = "--quiet"
	BusyVariable     = "--busy"
	RootSelector     = ":root"
	LayerSelector    = ":where(:root)"
)

const (
	BadgeClass         = "badge"
	BadgeQuietClass    = "quiet"
	BadgeRowClass      = "badge-row"
	BadgeColorVariable = "--badge-color"
)

const (
	ChipClass         = "chip"
	ChipToggleClass   = "toggle"
	ChipActiveClass   = "active"
	ChipColorVariable = "--chip-color"
)

const (
	CardClass          = "card"
	CardVisualClass    = "card-visual"
	CardFallbackClass  = "card-fallback"
	CardBodyClass      = "card-body"
	CardHeadClass      = "card-head"
	CardLinkClass      = "card-link"
	CardSubtitleClass  = "card-subtitle"
	CardMetaClass      = "card-meta"
	CardFooterClass    = "card-footer"
	CardAccentVariable = "--card-accent"
	CardArtClass       = "card-art"
	CardSideClass      = "card-side"
	CardFullClass      = "card-full"
	CardFaceClass      = "card-face"
	CardCompactClass   = "card-compact"
	CardHeroClass      = "card-hero"
	CardTaskClass      = "card-task"
)

const (
	MeterClass            = "meter"
	MeterRowClass         = "meter-row"
	MeterBarClass         = "meter-bar"
	MeterGaugeClass       = "meter-gauge"
	MeterGaugeFillClass   = "meter-gauge-fill"
	MeterIconClass        = "meter-icon"
	MeterLabelClass       = "meter-label"
	MeterValueClass       = "meter-value"
	MeterShareVariable    = "--meter-share"
	MeterColorVariable    = "--meter-color"
	MeterTrackVariable    = "--meter-track"
	MeterWidthVariable    = "--meter-width"
	MeterHeightVariable   = "--meter-height"
	ProgressColorVariable = "--pico-progress-color"
	ProgressTrackVariable = "--pico-progress-background-color"
)

const (
	FieldRowClass     = "field-row"
	FieldUnsetClass   = "field-unset"
	FieldLabelClass   = "field-label"
	FieldValueClass   = "field-value"
	FieldActionsClass = "field-actions"
	FieldNoteClass    = "field-note"
	FieldNotSet       = "not set"
	TooltipAttribute  = "data-tooltip"
	ValueAttribute    = "value"
)

const (
	EntryClass             = "entry"
	EntryInlineClass       = "inline"
	EntryRecordClass       = "record"
	EntryButtonsClass      = "entry-buttons"
	EntryAlternativeClass  = "entry-alternative"
	EntrySuggestionsSuffix = "-suggestions"
	EntryRevealInclude     = "closest form"
	EntryRevealTrigger     = "input changed delay:300ms"
	EntryLabelClass        = "entry-label"
	EntryTextKind          = "text"
	EntryNumberKind        = "number"
	EntryDateTimeKind      = "datetime-local"
	EntryDateKind          = "date"
	EntryLinkKind          = "url"
	EntryMailKind          = "email"
	EntryPasswordKind      = "password"
	EntrySearchKind        = "search"
	EntryResetScript       = "this.reset()"
	EntryRangeKind         = "range"
	EntryGroupKind         = "fieldset"
	EntryContentKind       = "content"
	EntryPreviewDelay      = "change delay:300ms"
	EntryCheckboxKind      = "checkbox"
	EntrySelectKind        = "select"
	EntryTextAreaKind      = "textarea"
	EntryFileKind          = "file"
	EntryCheckedValue      = "on"
)

const (
	EntityTagHeader    = "ETag"
	CacheControlHeader = "Cache-Control"
	RevalidateCache    = "no-cache"
)

const (
	FixtureEcho     = "echo"
	FixtureEchoPath = "/echo"
	FixtureTag      = "tag"
)
