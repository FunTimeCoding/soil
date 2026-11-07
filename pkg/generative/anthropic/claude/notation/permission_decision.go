package notation

type PermissionDecision struct {
	Decision   string `json:"decision"`
	Source     string `json:"source"`
	ReasonCode string `json:"reasonCode"`
	ReasonType string `json:"reasonType"`
	GrantScope string `json:"grantScope"`
}
