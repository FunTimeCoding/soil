package frame_probe

import "github.com/funtimecoding/go-library/pkg/chromium"

// FrameProbe answers the protocol questions in
// doc/ai/plan/gochromed-iframe-targets.md (devops/go-library):
// which fields the /json endpoint carries for iframe targets,
// what Target.getTargets reports, whether Page.getFrameTree in a
// tab session sees out-of-process frames, and whether an iframe
// target's identifier equals its frame identifier.
func FrameProbe() {
	c := chromium.NewEnvironment()
	printNotation()
	printTargets(c)
	printFrameTrees(c)
	printIframeRoots(c)
}
