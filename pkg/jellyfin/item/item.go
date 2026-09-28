package item

type Item struct {
	Identifier      string   `json:"id"`
	Name            string   `json:"name"`
	Type            string   `json:"type"`
	Overview        string   `json:"overview,omitempty"`
	Genres          []string `json:"genres,omitempty"`
	ProductionYear  int      `json:"production_year,omitzero"`
	CommunityRating float64  `json:"community_rating,omitzero"`
	OfficialRating  string   `json:"official_rating,omitempty"`
	RunTimeTicks    int64    `json:"run_time_ticks,omitzero"`
	SeriesName      string   `json:"series_name,omitempty"`
	SeasonNumber    int      `json:"season_number,omitzero"`
	EpisodeNumber   int      `json:"episode_number,omitzero"`
	IsFolder        bool     `json:"is_folder,omitzero"`
}
