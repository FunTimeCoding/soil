package mock_client

import "slices"

func (c *Client) membersAre(
	chat string,
	users []string,
) bool {
	var present []string

	for _, m := range c.chatMembers[chat] {
		present = append(present, m.User)
	}

	slices.Sort(present)
	wanted := slices.Clone(users)
	slices.Sort(wanted)

	return slices.Equal(present, wanted)
}
