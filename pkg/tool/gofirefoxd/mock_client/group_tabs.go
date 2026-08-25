package mock_client

func (c *Client) GroupTabs(
	tabIdentifiers []int,
	groupIdentifier int,
	title string,
	color string,
) (int, error) {
	identifier := groupIdentifier

	if identifier == 0 {
		c.groupIdentifier++
		identifier = c.groupIdentifier
	}

	c.groups[identifier] = &group{title: title, color: color}

	for _, tabIdentifier := range tabIdentifiers {
		for i, t := range c.tabs {
			if t.Identifier == tabIdentifier {
				c.tabs[i].GroupIdentifier = identifier
			}
		}
	}

	return identifier, nil
}
