package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gosourced",
	"Go source transformation daemon",
	"gosourced",
).WithInstructions(
	"Go source transformation daemon. AST-based rewriting for visibility changes, imports, package moves, and symbol renames. Multi-module - call list_modules and use_module before other operations.",
)

const (
	ListModules      = "list_modules"
	UseModule        = "use_module"
	ChangeVisibility = "change_visibility"
	RenameSymbol     = "rename_symbol"
	ExtractToFile    = "extract_to_file"
	AddImport        = "add_import"
	RemoveImport     = "remove_import"
)
