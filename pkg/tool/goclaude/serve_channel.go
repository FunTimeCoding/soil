package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/spf13/cobra"
	"time"
)

func serveChannel(
	c *command_context.Context,
	version string,
	r face.Reporter,
) *cobra.Command {
	var interval int
	result := &cobra.Command{
		Use:   "channel",
		Short: "Serve goclauded queue entries as channel events over stdio",
		Args:  cobra.NoArgs,
		Run: func(
			_ *cobra.Command,
			_ []string,
		) {
			s := channel.New(
				constant.Identity,
				version,
				constant.ChannelInstructions,
			)
			go runChannel(
				c,
				s,
				time.Duration(interval)*time.Second,
				recovery.New(logger.New(context.Background()), r),
			)
			s.Serve()
		},
	}
	result.Flags().IntVar(
		&interval,
		"interval",
		constant.ChannelInterval,
		"poll interval in seconds",
	)

	return result
}
