package shell_script

type Script struct {
	Text     string
	Printed  string
	Readable bool
	Parsed   bool
	Parent   *Script
	Calls    []*Call
	Scripts  []*Script
}
