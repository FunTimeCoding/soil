package constant

const (
	PointerClassLocator     PointerClass = "locator"
	PointerClassPlaceholder PointerClass = "placeholder"
	PointerClassSymbol      PointerClass = "symbol"
	PointerClassRoute       PointerClass = "route"
	PointerClassPath        PointerClass = "path"
	PointerClassImport      PointerClass = "import"
	PointerClassPattern     PointerClass = "pattern"
	PointerClassCommand     PointerClass = "command"
	PointerClassAbsolute    PointerClass = "absolute"
	PointerClassSystem      PointerClass = "system"
	PointerClassSibling     PointerClass = "sibling"
	PointerClassRepository  PointerClass = "repository"
	PointerClassShort       PointerClass = "short"
)

type PointerClass string
