package authenticate_command

type Command struct {
	Type  string `json:"type"`
	Token string `json:"access_token"`
}
