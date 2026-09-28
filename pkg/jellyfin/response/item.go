package response

type Item struct {
	Identifier        string   `json:"Id"`
	Name              string   `json:"Name"`
	Type              string   `json:"Type"`
	Overview          string   `json:"Overview"`
	Genres            []string `json:"Genres"`
	ProductionYear    int      `json:"ProductionYear"`
	CommunityRating   float64  `json:"CommunityRating"`
	OfficialRating    string   `json:"OfficialRating"`
	RunTimeTicks      int64    `json:"RunTimeTicks"`
	SeriesName        string   `json:"SeriesName"`
	ParentIndexNumber int      `json:"ParentIndexNumber"`
	IndexNumber       int      `json:"IndexNumber"`
	IsFolder          bool     `json:"IsFolder"`
}
