package frame_probe

import "github.com/funtimecoding/soil/pkg/browser/chromium"

func FrameProbe() {
	c := chromium.NewEnvironment()
	defer c.Close()
	printNotation()
	printTargets(c)
	printFrameTrees(c)
	printIframeRoots(c)
}
