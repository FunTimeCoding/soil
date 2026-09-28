package response

type NowPlaying struct {
	Identifier string `json:"Id"`
	Name       string `json:"Name"`
	Type       string `json:"Type"`
}
