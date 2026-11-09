package gw2

func GuildMembers(
	c *Client,
	tag string,
) []string {
	var result []string

	for _, guild := range c.Account().GuildLeader {
		if tag != "" && tag == c.Guild(guild).Tag {
			for _, member := range c.Members(guild) {
				result = append(result, member.Name)
			}
		}
	}

	return result
}
