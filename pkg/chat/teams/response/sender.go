package response

type Sender struct {
	User        *Identity `json:"user"`
	Application *Identity `json:"application"`
}
