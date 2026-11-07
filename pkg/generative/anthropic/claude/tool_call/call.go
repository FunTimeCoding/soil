package tool_call

type Call struct {
	Name               string
	Identifier         string
	Timestamp          string
	Directory          string
	Detail             string
	Input              string
	Result             string
	PermissionDecision string
	PermissionSource   string
	PermissionReason   string
}
