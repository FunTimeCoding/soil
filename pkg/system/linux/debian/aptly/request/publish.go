package request

type Publish struct {
	SourceKind    string          `json:"SourceKind"`
	Sources       []PublishSource `json:"Sources"`
	Architectures []string        `json:"Architectures"`
	Distribution  string          `json:"Distribution"`
	Signing       *SignOption     `json:"Signing,omitempty"`
}
