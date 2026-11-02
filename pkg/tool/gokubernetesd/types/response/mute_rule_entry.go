package response

type MuteRuleEntry struct {
	Identifier uint   `json:"identifier"`
	Reason     string `json:"reason"`
	Message    string `json:"message,omitempty"`
	Cluster    string `json:"cluster,omitempty"`
}
