package argument

type GetChannelHistory struct {
	Team    string `json:"team"`
	Channel string `json:"channel"`
	Limit   int    `json:"limit"`
}
