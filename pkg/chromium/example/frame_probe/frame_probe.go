package frame_probe

import "github.com/funtimecoding/soil/pkg/chromium"

func FrameProbe() {
	c := chromium.NewEnvironment()
	printNotation()
	printTargets(c)
	printFrameTrees(c)
	printIframeRoots(c)
}
