package client

import (
	"bufio"
	"errors"
	"fmt"
	panicErrors "github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"net"
	"syscall"
)

func Send(
	socketPath string,
	command string,
	arguments []string,
) (string, error) {
	connection, e := net.Dial("unix", socketPath)

	if e != nil {
		if errors.Is(e, syscall.ENOENT) || errors.Is(
			e,
			syscall.ECONNREFUSED,
		) {
			return "", fmt.Errorf(
				"no goprocessd instance running for this directory",
			)
		}

		return "", fmt.Errorf("connect: %w", e)
	}

	defer panicErrors.PanicClose(connection)
	line := command

	if len(arguments) > 0 {
		line = fmt.Sprintf("%s %s", command, join.Space(arguments...))
	}

	writer.Print(connection, "%s\n", line)
	scanner := bufio.NewScanner(connection)
	var lines []string

	for scanner.Scan() {
		text := scanner.Text()

		if text == "." {
			break
		}

		lines = append(lines, text)
	}

	if len(lines) == 0 {
		return "", fmt.Errorf("no response from server")
	}

	return join.NewLine(lines), nil
}
