package gitlab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"slices"
)

func (c *Client) Runners(all bool) ([]*runner.Runner, error) {
	var result []*gitlab.Runner
	var number int64

	for {
		var page []*gitlab.Runner
		var response *gitlab.Response
		var e error
		options := &gitlab.ListRunnersOptions{
			ListOptions: gitlab.ListOptions{
				PerPage: constant.PerPage100,
				Page:    number,
			},
		}

		if all {
			page, response, e = c.allRunners(options)
		} else {
			page, response, e = c.ownerRunners(options)
		}

		if e != nil {
			return nil, e
		}

		if false {
			// Paging does not make sense, duplicates are returned
			headerKeys := []string{
				"X-Next-Page",
				"X-Page",
				"X-Prev-Page",
				"X-Total",
				"X-Total-Pages",
				"X-Per-Page",
			}

			for k, v := range response.Header {
				if !slices.Contains(headerKeys, k) {
					continue
				}

				fmt.Printf("  %d Header %s: %v\n", number, k, v)
			}
		}

		result = append(result, page...)

		if int64(len(page)) < constant.PerPage100 {
			break
		}

		number++
	}

	return c.enrichRunners(runner.NewSlice(runner.Deduplicate(result))), nil
}
