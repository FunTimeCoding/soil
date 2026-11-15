package shell_script

func Parse(command string) *Script {
	result := &Script{Text: command, Printed: command, Readable: true}
	result.walk(result)

	return result
}
