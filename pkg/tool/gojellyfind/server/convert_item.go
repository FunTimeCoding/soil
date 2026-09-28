package server

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func convertItem(v *item.Item) *server.Item {
	result := &server.Item{Id: v.Identifier, Name: v.Name, Type: v.Type}

	if v.Overview != "" {
		result.Overview = new(v.Overview)
	}

	if len(v.Genres) > 0 {
		result.Genres = new(v.Genres)
	}

	if v.ProductionYear != 0 {
		result.ProductionYear = new(v.ProductionYear)
	}

	if v.CommunityRating != 0 {
		result.CommunityRating = new(float32(v.CommunityRating))
	}

	if v.OfficialRating != "" {
		result.OfficialRating = new(v.OfficialRating)
	}

	if v.RunTimeTicks != 0 {
		result.RunTimeTicks = new(v.RunTimeTicks)
	}

	if v.SeriesName != "" {
		result.SeriesName = new(v.SeriesName)
	}

	if v.SeasonNumber != 0 {
		result.SeasonNumber = new(v.SeasonNumber)
	}

	if v.EpisodeNumber != 0 {
		result.EpisodeNumber = new(v.EpisodeNumber)
	}

	if v.IsFolder {
		result.IsFolder = new(v.IsFolder)
	}

	return result
}
