package response

type SourceNatRule struct {
	Identifier      string `json:"uuid"`
	Enabled         Flag   `json:"enabled"`
	Interface       string `json:"interface"`
	Protocol        string `json:"protocol"`
	SourceNet       string `json:"source_net"`
	DestinationNet  string `json:"destination_net"`
	Target          string `json:"target"`
	TargetPort      string `json:"target_port"`
	Log             Flag   `json:"log"`
	Description     string `json:"description"`
}
