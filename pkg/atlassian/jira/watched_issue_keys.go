package jira

func (c *Client) WatchedIssueKeys() ([]string, error) {
	issues, e := c.SearchV3("issue IN watchedIssues() ORDER BY key ASC")

	if e != nil {
		return nil, e
	}

	var result []string

	for _, i := range issues {
		result = append(result, i.Key)
	}

	return result, nil
}
