package constant

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	system "github.com/funtimecoding/soil/pkg/system/constant"
)

const (
	ShebangPrefix          = "#!"
	EnvironmentInterpreter = "env"

	BlockOpenC          = "/*"
	BlockCloseC         = "*/"
	BlockOpenHypertext  = "<!--"
	BlockCloseHypertext = "-->"

	QuoteDouble   = `"`
	QuoteSingle   = "'"
	QuoteBacktick = "`"

	BinaryProbe = 8000

	UnplacedHeader = "Unplaced files:"
	SkippedHeader  = "Skipped files (binary or unreadable):"

	ByteOrderMark = "\xef\xbb\xbf"

	ColumnLanguage = "LANGUAGE"
	ColumnFiles    = "FILES"
	ColumnBlank    = "BLANK"
	ColumnComment  = "COMMENT"
	ColumnCode     = "CODE"
	ColumnPath     = "PATH"
	RowTotal       = "TOTAL"


	VendorDirectory      = "vendor"
	NodeModulesDirectory = "node_modules"
	FinderMetadataFile   = ".DS_Store"
)

var DefaultSkips = []string{
	constant.Directory,
	system.IdeaPath,
	FinderMetadataFile,
	VendorDirectory,
	NodeModulesDirectory,
}
