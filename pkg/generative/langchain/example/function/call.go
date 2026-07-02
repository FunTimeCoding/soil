package function

type Call struct {
	Tool  string         `json:"tool"`
	Input map[string]any `json:"tool_input"`
}
