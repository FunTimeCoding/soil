package mock_client

import "github.com/funtimecoding/soil/pkg/chat/teams/member"

func (c *Client) AddChatMember(
	chat string,
	m *member.Member,
) {
	c.chatMembers[chat] = append(c.chatMembers[chat], m)
}
