package mock_client

import "github.com/funtimecoding/soil/pkg/chat/teams/member"

func (c *Client) AddChannelMember(
	team string,
	channel string,
	m *member.Member,
) {
	k := key(team, channel)
	c.chatMembers[k] = append(c.chatMembers[k], m)
}
