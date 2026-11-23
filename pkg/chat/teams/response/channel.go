package response

type Channel struct {
	Identifier     string `json:"id"`
	DisplayName    string `json:"displayName"`
	Description    string `json:"description"`
	MembershipType string `json:"membershipType"`
	WebLocator     string `json:"webUrl"`
}
