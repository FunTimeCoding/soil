package main

import (
	brave "github.com/funtimecoding/soil/pkg/brave/example"
	chromium "github.com/funtimecoding/soil/pkg/chromium/example"
	"github.com/funtimecoding/soil/pkg/chromium/example/frame_probe"
	anthropic "github.com/funtimecoding/soil/pkg/generative/anthropic/site/example"
)

func main() {
	anthropic.Dump()

	if false {
		frame_probe.FrameProbe()
		chromium.Tab()
		brave.BookmarkSearch()
		brave.BookmarkNode()
		brave.BookmarkFile()
		brave.Extract()
		brave.Open()
		brave.Send()
		brave.Profile()
		chromium.Tabs()
		chromium.Tab()
		chromium.Open()
	}
}
