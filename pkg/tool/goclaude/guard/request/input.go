package request

type Input struct {
	SessionIdentifier string    `json:"session_id"`
	WorkingDirectory  string    `json:"cwd"`
	ToolName          string    `json:"tool_name"`
	ToolInput         ToolInput `json:"tool_input"`
}
