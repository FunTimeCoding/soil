package item

func New(
	identifier string,
	name string,
	itemType string,
	overview string,
	genres []string,
	productionYear int,
	communityRating float64,
	officialRating string,
	runTimeTicks int64,
	seriesName string,
	seasonNumber int,
	episodeNumber int,
	isFolder bool,
) *Item {
	return &Item{
		Identifier:      identifier,
		Name:            name,
		Type:            itemType,
		Overview:        overview,
		Genres:          genres,
		ProductionYear:  productionYear,
		CommunityRating: communityRating,
		OfficialRating:  officialRating,
		RunTimeTicks:    runTimeTicks,
		SeriesName:      seriesName,
		SeasonNumber:    seasonNumber,
		EpisodeNumber:   episodeNumber,
		IsFolder:        isFolder,
	}
}
