package constant

const (
	ReasonExternal   Reason = "external"
	ReasonHome       Reason = "home"
	ReasonImage      Reason = "image"
	ReasonImport     Reason = "import"
	ReasonNetwork    Reason = "network"
	ReasonPattern    Reason = "pattern"
	ReasonRoute      Reason = "route"
	ReasonSlash      Reason = "slash"
	ReasonSystem     Reason = "system"
	ReasonUnanchored Reason = "unanchored"
	ReasonUnknown    Reason = "unknown"
)

type Reason string
