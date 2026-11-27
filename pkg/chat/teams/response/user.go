package response

type User struct {
	Identifier  string `json:"id"`
	DisplayName string `json:"displayName"`
	Mail        string `json:"mail"`
	Principal   string `json:"userPrincipalName"`
	JobTitle    string `json:"jobTitle"`
}
