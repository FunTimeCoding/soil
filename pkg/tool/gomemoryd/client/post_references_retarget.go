package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/lint/types/retarget_change"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/retarget_report"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostReferencesRetarget(
	x context.Context,
	moves *move_set.Set,
	dryRun bool,
) ([]*retarget_report.Report, error) {
	body := client.PostReferencesRetargetJSONRequestBody{
		Paths:   []client.PathMove{},
		Symbols: []client.SymbolMove{},
		DryRun:  dryRun,
	}

	for _, m := range moves.Paths {
		body.Paths = append(body.Paths, client.PathMove{From: m.From, To: m.To})
	}

	for _, m := range moves.Symbols {
		body.Symbols = append(
			body.Symbols,
			client.SymbolMove{
				Directory:    m.Directory,
				Name:         m.Name,
				NewDirectory: m.NewDirectory,
				NewName:      m.NewName,
			},
		)
	}

	r, e := c.client.PostReferencesRetargetWithResponse(x, body)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("retarget references", r.HTTPResponse, r.Body)
	}

	result := []*retarget_report.Report{}

	for _, m := range r.JSON200.Memories {
		var changes []*retarget_change.Change

		for _, change := range m.Changes {
			if change.Reason != "" {
				changes = append(
					changes,
					retarget_change.NewManual(
						change.Line,
						change.Before,
						change.Reason,
					),
				)

				continue
			}

			changes = append(
				changes,
				retarget_change.New(change.Line, change.Before, change.After),
			)
		}

		result = append(
			result,
			retarget_report.New(m.Identifier, m.Name, changes),
		)
	}

	return result, nil
}
