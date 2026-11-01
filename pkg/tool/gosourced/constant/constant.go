package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"regexp"
)

var Identity = identity.New(
	"gosourced",
	"Go source transformation daemon",
	"gosourced",
).WithInstructions(
	"Go source transformation daemon. AST-based rewriting for visibility changes, imports, package moves, and symbol renames. Multi-module - call list_modules and use_module before other operations.",
)

const (
	HostEnvironment     = "GOSOURCE_HOST"
	PortEnvironment     = "GOSOURCE_PORT"
	InsecureEnvironment = "GOSOURCE_INSECURE"
	TokenEnvironment    = "GOSOURCE_TOKEN"
)

const (
	ListModules          = "list_modules"
	UseModule            = "use_module"
	FindReferences       = "find_references"
	MatchPattern         = "match_pattern"
	FindLiterals         = "find_literals"
	IntroduceConstructor = "introduce_constructor"
	ApplyPattern         = "apply_pattern"
	ListCalls            = "list_calls"
	ChangeVisibility     = "change_visibility"
	RenameSymbol         = "rename_symbol"
	MoveSymbol           = "move_symbol"
	MoveSymbols          = "move_symbols"
	ExtractType          = "extract_type"
	MovePackage          = "move_package"
	RenamePackage        = "rename_package"
	RenamePackageClause  = "rename_package_clause"
	RenameModule         = "rename_module"
	ExtractToFile        = "extract_to_file"
	AddImport            = "add_import"
	RemoveImport         = "remove_import"
	RemoveParameters     = "remove_parameters"
)

const (
	ConcernValidation      = "validation"
	ConcernRemoved         = "removed"
	ConcernValueUse        = "value_use"
	ConcernStillUsed       = "still_used"
	ConcernSideEffect      = "side_effect"
	ConcernLocal           = "local"
	ConcernUnusedParameter = "unused_parameter"
	ConcernWrittenOnly     = "written_only"
	ConcernConcurrentWrite = "concurrent_write"
	ConcernRenamed         = "renamed"
	ConcernMoved           = "moved"
)

const (
	ReferenceLimit  = 25
	MemberSeparator = "."
	StandalonePath  = "standalone"
	ParameterModule = "module"
	ParameterType   = "type"
	FieldSeparator  = " "
	ModuleKind      = "module"
	ModuleQualifier = "module"
	MissingPackage  = "missing"
	ConcernMissing  = "missing"
	ConcernChanged  = "changed"
)

const (
	SymbolFunction = "function"
	SymbolMethod   = "method"
	SymbolType     = "type"
	SymbolConstant = "constant"
	SymbolVariable = "variable"
)

const (
	IotaName          = "iota"
	BlankName         = "_"
	DotName           = "."
	TestClauseSuffix  = "_test"
	VersionPrefix     = "v"
	ToolPrefix        = "go-"
	ToolSuffix        = "-go"
	IotaMarker        = " // iota "
	VariantSeparator  = "@"
	OrdinalSeparator  = "#"
	KeySeparator      = " "
)

const (
	ClassUnchanged    = "unchanged"
	ClassMoved        = "moved"
	ClassRenamed      = "renamed"
	ClassMovedChanged = "moved-changed"
	ClassSignature    = "signature"
	ClassBody         = "body"
	ClassAdded        = "added"
	ClassRemoved      = "removed"
)

var Classes = []string{
	ClassUnchanged,
	ClassMoved,
	ClassRenamed,
	ClassMovedChanged,
	ClassSignature,
	ClassBody,
	ClassAdded,
	ClassRemoved,
}

var QualifiedReference = regexp.MustCompile(
	`"(?:[^"\\]|\\.)*" \. [A-Za-z_][A-Za-z0-9_]*`,
)

const ReferenceSeparator = " . "
const (
	OperationWrite         = "write"
	OperationRemove        = "remove"
	OperationMakeDirectory = "make_directory"
	OperationRename        = "rename"
)
