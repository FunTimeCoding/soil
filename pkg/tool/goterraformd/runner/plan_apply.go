package runner

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/host_mapping"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/types/plan"
	"path/filepath"
)

func (r *Runner) planApply(
	record *run.Run,
	target string,
) error {
	directory := filepath.Join(r.clonePath, r.terraformPath)
	scratch := system.TemporaryDirectory(constant.PlanPattern)
	defer system.Remove(scratch)
	file := filepath.Join(scratch, constant.PlanFile)
	arguments := []string{
		r.executable,
		constant.PlanArgument,
		join.Equals([]string{constant.OutArgument, file}),
		"-no-color",
	}

	if target != "" {
		arguments = append(arguments, fmt.Sprintf("-target=%s", target))
	}

	c := r.terraform(directory, arguments...)
	record.Output = c.OutputString
	record.ErrorOutput = c.ErrorString

	if c.Error != nil {
		return c.Error
	}

	s := r.terraform(
		directory,
		r.executable,
		constant.ShowArgument,
		constant.NotationArgument,
		file,
	)

	if s.Error != nil {
		record.ErrorOutput = s.ErrorString

		return s.Error
	}

	p := plan.New()

	if e := notation.DecodeBytes([]byte(s.OutputString), p); e != nil {
		return e
	}

	a := r.terraform(
		directory,
		r.executable,
		constant.ApplyArgument,
		"-auto-approve",
		"-no-color",
		file,
	)
	record.Output = join.Empty(record.Output, a.OutputString)
	record.ErrorOutput = a.ErrorString

	if a.Error != nil {
		return a.Error
	}

	r.pending = change.New(
		host_mapping.Addresses(p),
		host_mapping.New(p, constant.HostAttribute).Events(p.ResourceChanges),
	)

	return nil
}
