package constant

const (
	PolicyPessimistic = "pessimistic"
	PolicyOptimistic  = "optimistic"
	PolicySkip        = "skip"

	UnknownPolicy = "unknown missing policy %q, use pessimistic, optimistic or skip"

	DefaultThreshold = 30.0

	TestCommand         = "test"
	ListCommand         = "list"
	ListFormat          = "-f"
	TestPackageTemplate = "{{if or .TestGoFiles .XTestGoFiles}}{{.ImportPath}}{{end}}"
	CompileOnly         = "-c"
	CoverFlag           = "-cover"
	TestBinarySuffix    = ".test"
	ProfileSuffix       = ".out"
	TestList            = "-test.list"
	TestRun             = "-test.run"
	TestCoverProfile    = "-test.coverprofile"
	AnyTest             = ".*"
	TestPrefix          = "Test"
	ExactPattern        = "^%s$"
	WorkPrefix          = "gocrap-"
	CoverProfile        = "-coverprofile"
	CoverPackage   = "-coverpkg"
	Tool           = "tool"
	Cover          = "cover"
	FunctionReport = "-func"
	AllPackages    = "./..."
	ProfileFile    = "cover.out"

	MutantKilled = "KILLED"
	MutantLived  = "LIVED"

	TotalPrefix = "total:"
	Percent     = "%"

	MarkAbove = "✗"
	MarkWarn  = "▲"
	MarkOkay  = "✓"

	ScoreFormat      = "%.2f"
	CoverageFormat   = "%.1f%%"
	DeltaFormat      = "%+.2f"
	SummaryFormat    = "%d/%d above threshold %.0f | combined %.2f | average %.2f\n"
	SkippedFormat    = "%d without coverage skipped\n"
	FailFormat       = "%d function(s) above threshold %.0f\n"
	BaselineFormat   = "combined delta against baseline %+.2f\n"
	RegressionHeader = "regressions:"
	RegressionFormat = "%d function(s) regressed against baseline\n"
	MatrixSummary    = "%d tests, %d functions covered, %d by a single test\n"
	ParallelRefused  = "attribution refused: parallel subtests would smear the matrix, narrow the pattern to exclude:\n%s\n"
	AttributeUnscoped = "attribution refused: ./... compiles every test binary against the whole module, name the packages to attribute\n"
	MissingSuffix    = " (missing)"
	UntrustedFormat  = " (%d/%d mutants killed)"

	DefaultTolerance = 0.01
	DeltaNew         = "new"
	DeltaSame        = "-"

	ColumnMark     = " "
	ColumnScore    = "CRAP"
	ColumnCC       = "CC"
	ColumnCoverage = "COVERAGE"
	ColumnDelta    = "DELTA"
	ColumnPrevious = "PREVIOUS"
	ColumnFunction = "FUNCTION"
	ColumnLocation = "LOCATION"
	ColumnAlone    = "ALONE"
	ColumnTotal    = "TOTAL"
	ColumnTest     = "TEST"
)
