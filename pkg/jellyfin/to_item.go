package jellyfin

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/jellyfin/response"
)

func toItem(r *response.Item) *item.Item {
	return item.New(
		r.Identifier,
		r.Name,
		r.Type,
		r.Overview,
		r.Genres,
		r.ProductionYear,
		r.CommunityRating,
		r.OfficialRating,
		r.RunTimeTicks,
		r.SeriesName,
		r.ParentIndexNumber,
		r.IndexNumber,
		r.IsFolder,
	)
}
