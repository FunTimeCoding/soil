package model_context

import "github.com/mattermost/mattermost/server/public/model"

func channelTypeName(t model.ChannelType) string {
	switch t {
	case model.ChannelTypeOpen:
		return "public"
	case model.ChannelTypePrivate:
		return "private"
	case model.ChannelTypeDirect:
		return "dm"
	case model.ChannelTypeGroup:
		return "group_dm"
	}

	return string(t)
}
