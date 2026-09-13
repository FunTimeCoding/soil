package constant

const (
	ExtendedGet          = "hx-get"
	ExtendedPost         = "hx-post"
	ExtendedTarget       = "hx-target"
	ExtendedSwap         = "hx-swap"
	ExtendedOutOfBand    = "hx-swap-oob"
	ExtendedTrigger      = "hx-trigger"
	ExtendedInclude      = "hx-include"
	ExtendedValue        = "hx-vals"
	ExtendedConfirm      = "hx-confirm"
	ExtendedExtension    = "hx-ext"
	ExtendedIndicator    = "hx-indicator"
	ExtendedOnPrefix     = "hx-on:"
	ExtendedAfterRequest = "hx-on::after-request"
	ServerSideConnect    = "sse-connect"
	ServerSideSwap       = "sse-swap"
)

const (
	SwapInner  = "innerHTML"
	SwapOuter  = "outerHTML"
	SwapAppend = "beforeend"
	SwapDelete = "delete"
	SwapNone   = "none"
)

const (
	ExtendedIndicatorClass = "htmx-indicator"
	IndicatorMarkClass     = "indicator-mark"
)

const (
	TriggerLoad   = "load"
	TriggerChange = "change"
	TriggerType   = "keyup changed delay:200ms"
)
