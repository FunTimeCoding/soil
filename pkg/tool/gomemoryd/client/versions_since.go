package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/version_entry"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *RestClient) VersionsSince(
	x context.Context,
	since string,
	limit int,
) ([]version_entry.Entry, error) {
	r, e := c.http.GetVersionsWithResponse(
		x,
		&client.GetVersionsParams{Since: since, Limit: &limit},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("versions", r.HTTPResponse, r.Body)
	}

	var result []version_entry.Entry

	for _, v := range *r.JSON200 {
		result = append(
			result,
			version_entry.Entry{
				MemoryIdentifier: v.MemoryIdentifier,
				Name:             v.Name,
				Description:      v.Description,
				ChangeType:       v.ChangeType,
				Source:           v.Source,
				ChangedAt:        v.ChangedAt,
			},
		)
	}

	return result, nil
}
