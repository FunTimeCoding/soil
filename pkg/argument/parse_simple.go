package argument

import "os"

func (i *Instance) ParseSimple() {
	i.exitOnParseError(i.flags.Parse(os.Args[1:]))
}
