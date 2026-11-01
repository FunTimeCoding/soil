package face

import "github.com/mattermost/mattermost/server/public/model"

type ChannelClient interface {
	AllChannels(
		limit int,
		offset int,
	) ([]*model.ChannelWithTeamData, error)
	Channel(identifier string) (*model.Channel, error)
	ChannelByName(
		t *model.Team,
		name string,
	) (*model.Channel, error)
	ChannelUsers(h *model.Channel) ([]*model.User, error)
	Channels(
		t *model.Team,
		u *model.User,
	) ([]*model.Channel, error)
	DefaultChannel() *model.Channel
	DefaultTeam() *model.Team
	DirectMessage(
		u *model.User,
		text string,
	) (*model.Post, error)
	MustAllChannels(
		limit int,
		offset int,
	) []*model.ChannelWithTeamData
	MustChannel(identifier string) *model.Channel
	MustChannelByName(
		t *model.Team,
		name string,
	) *model.Channel
	MustChannelUsers(h *model.Channel) []*model.User
	MustChannels(
		t *model.Team,
		u *model.User,
	) []*model.Channel
	MustDirectMessage(
		u *model.User,
		text string,
	) *model.Post
	MustTeam(name string) *model.Team
	MustTeamChannel(name string) *model.Channel
	MustTeams(userIdentifier string) []*model.Team
	MustUpdateChannelHeader(
		h *model.Channel,
		prefix string,
		text string,
		separator string,
	) *model.Channel
	SearchChannels(
		t *model.Team,
		query string,
	) ([]*model.Channel, error)
	Team(name string) (*model.Team, error)
	TeamChannel(name string) (*model.Channel, error)
	Teams(userIdentifier string) ([]*model.Team, error)
	UpdateChannelHeader(
		h *model.Channel,
		prefix string,
		text string,
		separator string,
	) (*model.Channel, error)
}
