package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *RestClient) VersionsSince(
	since string,
	limit int,
) []VersionEntry {
	r, e := c.http.GetVersions(
		context.Background(),
		&client.GetVersionsParams{
			Since: since,
			Limit: &limit,
		},
	)

	if e != nil {
		return nil
	}

	defer errors.PanicClose(r.Body)
	parsed, e := client.ParseGetVersionsResponse(r)

	if e != nil || parsed.JSON200 == nil {
		return nil
	}

	var result []VersionEntry

	for _, v := range *parsed.JSON200 {
		result = append(
			result,
			VersionEntry{
				MemoryIdentifier: v.MemoryIdentifier,
				Name:             v.Name,
				Description:      v.Description,
				ChangeType:       v.ChangeType,
				Source:           v.Source,
				ChangedAt:        v.ChangedAt,
			},
		)
	}

	return result
}
