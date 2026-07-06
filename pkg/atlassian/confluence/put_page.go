package confluence

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page/page_put"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) PutPage(
	identifier string,
	title string,
	body string,
	version int,
	message string,
	status string,
) (*page.Page, error) {
	result, e := c.basic.PutV2Path(
		fmt.Sprintf("%s/%s", constant.Page, identifier),
		page_put.NewWithStatus(
			identifier,
			title,
			body,
			version,
			message,
			status,
		).Encode(),
	)

	if e != nil {
		return nil, e
	}

	var p *response.Page
	notation.MustDecode(result, &p, false)

	return page.New(p, c.host), nil
}
