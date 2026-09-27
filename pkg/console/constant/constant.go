package constant

import (
	"github.com/fatih/color"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

const (
	GreenColor  = "green"
	RedColor    = "red"
	YellowColor = "yellow"

	// Output formats
	FormatText     = "text"
	FormatNotation = "notation"
	FormatMarkdown = "markdown"

	TagAge         = "age"
	TagAssignee    = "assignee"
	TagCategory    = "category"
	TagChanges     = "changes"
	TagCluster     = "cluster"
	TagComment     = "comment"
	TagConcerns    = "concerns"
	TagCopyable    = "copyable"
	TagDense       = "dense"
	TagDescription = "description"
	TagEmoji       = "emoji"
	TagFilter      = "filter"
	TagFingerprint = "fingerprint"
	TagGraph       = "graph"
	TagHost        = "host"
	TagIdentifier  = "identifier"
	TagInstance    = "instance"
	TagInvestigate = "investigate"
	TagKey         = "key"
	TagLabels      = "labels"
	TagMarkdown    = "markdown"
	TagName        = "name"
	TagProject     = "project"
	TagRunbook     = "runbook"
	TagScore       = "score"
	TagState       = "state"
	TagStatus      = "status"
	TagTimestamp   = "timestamp"
	TagType        = "type"
	TagUsage       = "usage"
	TagWiki        = "wiki"
)

var (
	Blue    = color.New(color.FgBlue).SprintfFunc()
	Cyan    = color.New(color.FgCyan).SprintfFunc()
	Green   = color.New(color.FgGreen).SprintfFunc()
	Magenta = color.New(color.FgMagenta).SprintfFunc()
	Red     = color.New(color.FgRed).SprintfFunc()
	Yellow  = color.New(color.FgYellow).SprintfFunc()

	ColorFormat         = option.New().Color()
	ExtendedColorFormat = option.New().Extended().Color()
)
