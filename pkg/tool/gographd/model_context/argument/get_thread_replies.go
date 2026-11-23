package argument

type GetThreadReplies struct {
	Team    string `json:"team"`
	Channel string `json:"channel"`
	Message string `json:"message"`
	Limit   int    `json:"limit"`
}
