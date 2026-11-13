package constant

const (
	SectionClass             = "section"
	ActionClass              = "action"
	ActionSmallClass         = "small"
	NumericColumnClassPrefix = "numeric-"
	AbsentClass              = "absent"
	UnsetMark                = "—"
	ClusterClass             = "cluster"
	HeadingRowClass          = "heading-row"
	ActionRowClass           = "action-row"
	FactsClass               = "facts"
	RowScope                 = "row"
	GlanceClass              = "glance"
	GlanceColumnsClass       = "glance-columns"
	GlanceColumnsVariable    = "--glance-columns"
	GlanceGroupClass         = "glance-group"
	GlanceLabelClass         = "glance-label"
	GlanceValueClass         = "glance-value"
	TileClass                = "tile"
	TileGridClass            = "tile-grid"
	DialClass                = "dial"
	DialCenterClass          = "dial-center"
	DialNeedleClass          = "dial-needle"
	DialPointClass           = "dial-point"
	DialTickClass            = "dial-tick"
	DialAngleVariable        = "--dial-angle"
	DialBackgroundVariable   = "--dial-background"
	StatusDotClass           = "status-dot"
	PagerClass               = "pager"
	PagerPositionClass       = "pager-position"
	SearchAttribute          = "data-search"
	BodyTarget               = "body"
)

const (
	NumericColumnLimit = 12
	DialDegrees        = 360.0
	DialQuarter        = 90.0
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
	CardClampClass     = "card-clamp"
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
	MeterIconClass        = "meter-icon"
	MeterLabelClass       = "meter-label"
	MeterValueClass       = "meter-value"
	MeterWidthVariable    = "--meter-width"
	MeterHeightVariable   = "--meter-height"
	ProgressColorVariable = "--pico-progress-color"
	ProgressTrackVariable = "--pico-progress-background-color"
)

const (
	PlateClass         = "plate"
	PlateGroupClass    = "plate-group"
	PlateNameClass     = "plate-name"
	PlateHeadlineClass = "plate-headline"
	PlateCaptionClass  = "plate-caption"
	PlateRowClass      = "plate-row"
	PlateLabelClass    = "plate-label"
	PlateValueClass    = "plate-value"
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
	OpeningRowClass     = "opening-row"
	OpeningDetailClass  = "opening-detail"
	OpeningContentClass = "opening-content"
	OpeningDetailSuffix = "-detail"
	ExpandedAttribute   = "aria-expanded"
	CurrentAttribute    = "aria-current"
	HighlightClass      = "highlight"
)

const (
	EntryClass             = "entry"
	EntryInlineClass       = "inline"
	EntrySidewaysClass     = "sideways"
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
