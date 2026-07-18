package guard

type input struct {
	ToolName  string    `json:"tool_name"`
	ToolInput toolInput `json:"tool_input"`
}
