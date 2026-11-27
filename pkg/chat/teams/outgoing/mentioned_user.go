package outgoing

type MentionedUser struct {
	Identifier   string `json:"id"`
	DisplayName  string `json:"displayName"`
	IdentityType string `json:"userIdentityType"`
}
