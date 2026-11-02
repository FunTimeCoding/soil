package gitlab

import "context"

func (c *Client) CompareFiles(
	x context.Context,
	project int64,
	from string,
	to string,
) ([]string, error) {
	result, e := c.Compare(x, project, from, to)

	if e != nil {
		return nil, e
	}

	files := make([]string, 0, len(result.Diffs))

	for _, d := range result.Diffs {
		files = append(files, d.NewPath)
	}

	return files, nil
}
