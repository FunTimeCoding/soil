package unit

import (
	"bufio"
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	generative "github.com/funtimecoding/soil/pkg/generative/model_context/server"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/mark3labs/mcp-go/server"
	"io"
	"testing"
)

func TestModelContextServerLegacyProtocol(t *testing.T) {
	s := server.NewStdioServer(
		server.NewMCPServer("probe", library.DefaultVersion),
	)
	s.SetContextFunc(generative.LegacyProtocol)
	inputReader, inputWriter := io.Pipe()
	outputReader, outputWriter := io.Pipe()
	done := make(chan error)
	go func() {
		done <- s.Listen(context.Background(), inputReader, outputWriter)
	}()
	_, e := fmt.Fprintln(inputWriter, constant.ModernDiscoverRequest)
	assert.FatalOnError(t, e)
	line, e := bufio.NewReader(outputReader).ReadString('\n')
	assert.FatalOnError(t, e)
	var r DiscoverResponse
	notation.MustDecode(line, &r, true)
	errors.PanicClose(inputWriter)
	assert.FatalOnError(t, <-done)
	assert.Strings(
		t,
		[]string{"2025-11-25", "2025-06-18", "2025-03-26", "2024-11-05"},
		r.Result.SupportedVersions,
	)
}
