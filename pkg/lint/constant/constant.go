package constant

const (
	FrontMatterDelimiter = "---"

	FrontMatterDelimiterKey  = "front_matter_delimiter"
	FrontMatterDelimiterText = "No front matter delimiter"

	ImportBlankKey  = "import_blank"
	ImportBlankText = "Import block contains blank line"

	UnsortedImportsKey  = "unsorted_imports"
	UnsortedImportsText = "Imports are not sorted"

	SingleMultiImportKey  = "single_multi_import"
	SingleMultiImportText = "Single multi import"

	EmptyFunctionBodyKey  = "empty-function-body"
	EmptyFunctionBodyText = "Function body with only whitespace"

	ErrorVariableKey  = "err_variable"
	ErrorVariableText = "Use e instead of err for error variable"

	PackageNameKey  = "package_name"
	PackageNameText = "Blacklisted package name"

	MissingTestFileKey  = "missing_test_file"
	MissingTestFileText = "No test file in package"

	MissingBlankBeforeControlKey  = "missing_blank_before_control"
	MissingBlankBeforeControlText = "Missing blank line before control block"

	MissingBlankBeforeExitKey  = "missing_blank_before_exit"
	MissingBlankBeforeExitText = "Missing blank line before return/break/continue"

	ExtraneousBlankLineKey  = "extraneous_blank_line"
	ExtraneousBlankLineText = "Extraneous blank line"

	MissingBlankAfterControlKey  = "missing_blank_after_control"
	MissingBlankAfterControlText = "Missing blank line after control block"

	MissingSentryKey  = "missing_sentry"
	MissingSentryText = "No sentry reporter in program"

	StrayConstantKey  = "stray_const"
	StrayConstantText = "Top-level const outside of a constant file or constant package"

	BlankInsideFunctionKey  = "blank_inside_function"
	BlankInsideFunctionText = "Blank line between statements inside function body"

	MissingBlankBeforeDeclarationKey  = "missing_blank_before_declaration"
	MissingBlankBeforeDeclarationText = "Missing blank line before declaration"

	ExtraneousTopLevelBlankKey  = "extraneous_top_level_blank"
	ExtraneousTopLevelBlankText = "Extraneous blank line between top-level declarations"

	MissingBlankBetweenVariableConstantKey  = "missing_blank_between_var_const"
	MissingBlankBetweenVariableConstantText = "Missing blank line between top-level const and var declarations"

	StrayBinaryKey  = "stray_binary"
	StrayBinaryText = "Stray binary detected"

	VariableGroupingKey  = "var_grouping"
	VariableGroupingText = "Consecutive var declarations should use var (...) block"

	DeadPointerKey  = "dead_pointer"
	DeadPointerText = "Referenced path does not exist"

	AbsolutePointerKey  = "absolute_pointer"
	AbsolutePointerText = "Absolute path reference"
)

var PackageBlocklist = []string{"api"}
