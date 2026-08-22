package runner

import (
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"path/filepath"
)

func (r *Runner) terraformInit() {
	directory := filepath.Join(r.clonePath, r.terraformPath)
	r.logger.Structured("terraform_init")
	c := r.newRun().NoPanic()
	c.Directory = directory
	c.Start(constant.Command, "init", "-json")

	if c.Error != nil && r.needsUpgrade(c.OutputString) {
		r.logger.Structured("terraform_init_upgrade")
		u := r.newRun()
		u.Directory = directory
		u.Start(constant.Command, "init", "-upgrade")

		return
	}

	if c.Error != nil {
		panic(c.Error)
	}
}
