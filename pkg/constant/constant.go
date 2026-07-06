package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text/option"
)

const (
	Go      = "go"
	Mod     = "mod"
	Tidy    = "tidy"
	Edit    = "edit"
	Build   = "build"
	Get     = "get"
	Version = "version"

	LinkerFlagsArgument = "-ldflags"
	OutputArgument      = "-o"
	TagsArgument        = "-tags"

	VersionArgument = "-go"

	LinkerSetVariable = "-X"

	NativeEnabled = "CGO_ENABLED"
	System        = "GOOS"
	Architecture  = "GOARCH"
	Proxy         = "GOPROXY"

	Direct = "direct"

	HomeEnvironment = "HOME"

	CurrentDirectory = "."

	GoExtension          = ".go"
	GraphicExtension     = ".png"
	HypertextExtension   = ".html"
	MarkdownExtension    = ".md"
	MarkupExtension      = ".yaml"
	NotationLogExtension = ".jsonl"
	ShortMarkupExtension = ".yml"

	LatestVersion = "latest"

	GeneratedFile = "generated.go"

	PhysicalTest0 = "00:00:00:00:00:00"
	PhysicalTest1 = "00:00:00:00:00:01"
	PhysicalTest2 = "00:00:00:00:00:02"

	Unauthorized = "Unauthorized"

	InvalidRequestBody = "invalid request body"
	UnexpectedError    = "unexpected error"

	SoilModule = "github.com/funtimecoding/soil"

	DefaultVersion = "1.0.0"

	TestClient = "test-client"

	TestDatabase = "test.db"

	TestSuffix = "_test.go"
)

// For console status option
const (
	LabelKey = "label"
	TagKey   = "tag"
)

var (
	StartOfTime       = assert.NewDay(1)
	CompactWhitespace = option.Compact()
)
