package response

type Rule struct {
	Identifier      string `json:"uuid"`
	Enabled         Flag   `json:"enabled"`
	Sequence        string `json:"sequence"`
	Interface       string `json:"interface"`
	Direction       string `json:"direction"`
	Action          string `json:"action"`
	Protocol        string `json:"protocol"`
	SourceNet       string `json:"source_net"`
	SourcePort      string `json:"source_port"`
	DestinationNet  string `json:"destination_net"`
	DestinationPort string `json:"destination_port"`
	Log             Flag   `json:"log"`
	Automatic       Flag   `json:"is_automatic"`
	Description     string `json:"description"`
	Categories      string `json:"categories"`
}
