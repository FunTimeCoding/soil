package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
	"github.com/funtimecoding/soil/pkg/chat/teams/mock_client"
	"github.com/funtimecoding/soil/pkg/chat/unit/teams_tester"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/microsoft/graph_server"
	"net/http"
	"testing"
)

func TestTeamsClientAndMockShareFace(t *testing.T) {
	faces := []face.TeamsClient{
		teams_tester.New(t, graph_server.New(t)),
		mock_client.New(t),
	}
	assert.Count(t, 2, faces)
}

func TestJoinedTeamsMapped(t *testing.T) {
	s := graph_server.New(t).Route(
		constant.TeamsPath,
		http.StatusOK,
		`{"value":[{"id":"team-1","displayName":"Security","description":"CVE work"}]}`,
	)
	result, e := teams_tester.New(t, s).Teams()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, result)
	assert.String(t, "team-1", result[0].Identifier)
	assert.String(t, "Security", result[0].Name)
	assert.String(t, "CVE work", result[0].Description)
}

func TestChannelMessagesKeepUserMessagesOnly(t *testing.T) {
	path := fmt.Sprintf(
		constant.TeamsChannelMessagesPath,
		"team-1",
		"19:channel@thread.tacv2",
	)
	s := graph_server.New(t).Route(
		path,
		http.StatusOK,
		`{"value":[
			{"id":"1","messageType":"message","createdDateTime":"2026-11-23T10:00:00Z","from":{"user":{"displayName":"November"}},"body":{"contentType":"html","content":"<p>patched?</p><p>all hosts</p>"},
			 "replies":[
				{"id":"2","messageType":"message","from":{"user":{"displayName":"Juliett"}},"body":{"contentType":"text","content":"yes"}},
				{"id":"3","messageType":"systemEventMessage","body":{"contentType":"html","content":"<systemEventMessage/>"}}
			 ]},
			{"id":"4","messageType":"systemEventMessage","body":{"contentType":"html","content":"<systemEventMessage/>"}},
			{"id":"5","messageType":"message","deletedDateTime":"2026-11-23T11:00:00Z","body":{"contentType":"text","content":""}}
		]}`,
	)
	result, e := teams_tester.New(t, s).ChannelMessages(
		"team-1",
		"19:channel@thread.tacv2",
		0,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, result)
	assert.String(t, "1", result[0].Identifier)
	assert.String(t, "November", result[0].Author)
	assert.String(t, "patched?\nall hosts", result[0].Text)
	assert.Integer(t, 1, result[0].Replies)
	x := s.Exchanges(path)
	assert.Count(t, 1, x)
	assert.String(t, "20", x[0].Query.Get("$top"))
	assert.String(t, "replies", x[0].Query.Get("$expand"))
}

func TestChatsNewestFirstWithPageSize(t *testing.T) {
	s := graph_server.New(t).Route(
		constant.TeamsChatsPath,
		http.StatusOK,
		`{"value":[{"id":"19:a@thread.v2","topic":"Squad","chatType":"group","lastUpdatedDateTime":"2026-11-23T10:00:00Z"},{"id":"19:b@thread.v2","chatType":"oneOnOne"}]}`,
	)
	result, e := teams_tester.New(t, s).Chats(1)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, result)
	assert.String(t, "Squad", result[0].Topic)
	assert.String(t, "group", result[0].Kind)
	x := s.Exchanges(constant.TeamsChatsPath)
	assert.String(t, "1", x[0].Query.Get("$top"))
	assert.String(
		t,
		"lastMessagePreview/createdDateTime desc",
		x[0].Query.Get("$orderby"),
	)
}

func TestChatMessagesPageSizeCapped(t *testing.T) {
	path := fmt.Sprintf(constant.TeamsChatMessagesPath, "19:a@thread.v2")
	s := graph_server.New(t).Route(
		path,
		http.StatusOK,
		`{"value":[{"id":"7","messageType":"message","from":{"application":{"displayName":"Bot"}},"body":{"contentType":"text","content":"<b>raw</b>"}}]}`,
	)
	result, e := teams_tester.New(t, s).ChatMessages("19:a@thread.v2", 500)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, result)
	assert.String(t, "Bot", result[0].Author)
	assert.String(t, "<b>raw</b>", result[0].Text)
	x := s.Exchanges(path)
	assert.String(t, "50", x[0].Query.Get("$top"))
	assert.String(t, "createdDateTime desc", x[0].Query.Get("$orderby"))
}

func TestChatMessageSentAsHypertextWithLineBreaks(t *testing.T) {
	path := fmt.Sprintf(constant.TeamsChatMessagesPath, "19:a@thread.v2")
	s := graph_server.New(t).Route(
		path,
		http.StatusCreated,
		`{"id":"9","messageType":"message","from":{"user":{"displayName":"Juliett"}},"body":{"contentType":"html","content":"<p>alfa &lt;b&gt;</p>"}}`,
	)
	result, e := teams_tester.New(t, s).SendChatMessage(
		"19:a@thread.v2",
		"alfa <b>\nbravo",
		nil,
	)
	assert.FatalOnError(t, e)
	assert.String(t, "9", result.Identifier)
	assert.String(t, "Juliett", result.Author)
	x := s.Exchanges(path)
	assert.Count(t, 1, x)
	assert.String(t, http.MethodPost, x[0].Method)
	assert.String(
		t,
		`{"body":{"contentType":"html","content":"alfa &lt;b&gt;<br>bravo"}}`,
		x[0].Body,
	)
}

func TestChatMembersMapped(t *testing.T) {
	path := fmt.Sprintf(constant.TeamsChatMembersPath, "19:a@thread.v2")
	s := graph_server.New(t).Route(
		path,
		http.StatusOK,
		`{"value":[{"id":"membership-1","displayName":"Juliett","userId":"user-9"}]}`,
	)
	result, e := teams_tester.New(t, s).ChatMembers("19:a@thread.v2")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, result)
	assert.String(t, "Juliett", result[0].Name)
	assert.String(t, "user-9", result[0].User)
}

func TestChatMessageTagsMentionedMember(t *testing.T) {
	path := fmt.Sprintf(constant.TeamsChatMessagesPath, "19:a@thread.v2")
	s := graph_server.New(t).Route(
		path,
		http.StatusCreated,
		`{"id":"9","messageType":"message","body":{"contentType":"html","content":"<at id=\"0\">Juliett</at> busy?"}}`,
	)
	_, e := teams_tester.New(t, s).SendChatMessage(
		"19:a@thread.v2",
		"@Juliett busy?",
		[]*member.Member{member.New("Juliett", "user-9")},
	)
	assert.FatalOnError(t, e)
	assert.String(
		t,
		`{"body":{"contentType":"html","content":"<at id=\"0\">Juliett</at> busy?"},"mentions":[{"id":0,"mentionText":"Juliett","mentioned":{"user":{"id":"user-9","displayName":"Juliett","userIdentityType":"aadUser"}}}]}`,
		s.Exchanges(path)[0].Body,
	)
}

func TestChatMessageWithoutNamedMentionInvalid(t *testing.T) {
	s := graph_server.New(t)
	_, e := teams_tester.New(t, s).SendChatMessage(
		"19:a@thread.v2",
		"busy?",
		[]*member.Member{member.New("Juliett", "user-9")},
	)
	assert.True(t, validation.Is(e))
	assert.Count(
		t,
		0,
		s.Exchanges(
			fmt.Sprintf(constant.TeamsChatMessagesPath, "19:a@thread.v2"),
		),
	)
}

func TestChannelsAndRepliesMapped(t *testing.T) {
	s := graph_server.New(t).
		Route(
			fmt.Sprintf(constant.TeamsChannelsPath, "team-1"),
			http.StatusOK,
			`{"value":[{"id":"19:c@thread.tacv2","displayName":"General","membershipType":"standard","webUrl":"https://teams.example/c"}]}`,
		).
		Route(
			fmt.Sprintf(
				constant.TeamsRepliesPath,
				"team-1",
				"19:c@thread.tacv2",
				"100",
			),
			http.StatusOK,
			`{"value":[{"id":"101","messageType":"message","from":{"user":{"displayName":"Kilo"}},"body":{"contentType":"text","content":"ok"}}]}`,
		)
	c := teams_tester.New(t, s)
	channels, e := c.Channels("team-1")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, channels)
	assert.String(t, "General", channels[0].Name)
	assert.String(t, "standard", channels[0].Membership)
	assert.String(t, "https://teams.example/c", channels[0].Link)
	replies, f := c.Replies("team-1", "19:c@thread.tacv2", "100", 0)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, replies)
	assert.String(t, "Kilo", replies[0].Author)
}
