package main

import (
	discord "github.com/funtimecoding/soil/pkg/chat/discord/example"
	mattermost "github.com/funtimecoding/soil/pkg/chat/mattermost/example"
	telegram "github.com/funtimecoding/soil/pkg/chat/telegram/example"
)

func main() {
	mattermost.Since()

	if false {
		mattermost.DeleteTwice()
		mattermost.Paging()
		mattermost.Latest()
		mattermost.Before()
		discord.DeleteLoop()
		mattermost.Team()
		mattermost.Dialog()
		mattermost.Support()
		mattermost.Post()
		telegram.OllamaSession()
		telegram.OllamaReply()
		telegram.Update()
		telegram.Echo()
		telegram.User()
	}
}
