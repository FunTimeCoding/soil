package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/board_limit"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
)

func (c *Client) BoardLimits(identifier int) ([]*board_limit.Limit, error) {
	b, e := c.Board(identifier)

	if e != nil {
		return nil, e
	}

	o, f := c.BoardConfiguration(identifier)

	if f != nil {
		return nil, f
	}

	if o.ColumnConfig.ConstraintType != constant.IssueCountType {
		return nil, fmt.Errorf(
			"unexpected constraint type: %s",
			o.ColumnConfig.ConstraintType,
		)
	}

	var result []*board_limit.Limit

	for _, column := range o.ColumnConfig.Columns {
		if column.Min == 0 && column.Max == 0 {
			continue
		}

		result = append(
			result,
			board_limit.New(b.Name, column.Name, column.Max),
		)
	}

	return result, nil
}
