package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/spf13/cobra"
	"testing"
)

func TestANestedCommandIsRecordedByItsFullPath(t *testing.T) {
	root := &cobra.Command{Use: "gotest"}
	group := &cobra.Command{Use: "character"}
	group.AddCommand(
		&cobra.Command{
			Use: constant.List,
			Run: func(*cobra.Command, []string) {},
		},
	)
	root.AddCommand(group)
	assert.Strings(
		t,
		[]string{"begin gotest character list", "record gotest character list"},
		runInstrumented(t, root, "character", constant.List).entries,
	)
}

func TestTheRootsOwnPreRunStillRuns(t *testing.T) {
	ran := false
	root := &cobra.Command{
		Use: "gotest",
		PersistentPreRun: func(*cobra.Command, []string) {
			ran = true
		},
	}
	root.AddCommand(
		&cobra.Command{
			Use: constant.List,
			Run: func(*cobra.Command, []string) {},
		},
	)
	runInstrumented(t, root, constant.List)
	assert.True(t, ran)
}
