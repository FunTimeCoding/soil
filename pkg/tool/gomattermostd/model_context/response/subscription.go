package response

type Subscription struct {
	Root      string `json:"root"`
	Label     string `json:"label"`
	Channel   string `json:"channel"`
	LastEvent string `json:"last_event"`
}
