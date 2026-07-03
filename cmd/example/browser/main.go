package main

import (
	brave "github.com/funtimecoding/go-library/pkg/brave/example"
	chromium "github.com/funtimecoding/go-library/pkg/chromium/example"
	"github.com/funtimecoding/go-library/pkg/chromium/example/frame_probe"
	anthropic "github.com/funtimecoding/go-library/pkg/generative/anthropic/site/example"
)

func main() {
	frame_probe.FrameProbe()

	if false {
		anthropic.Dump()
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
