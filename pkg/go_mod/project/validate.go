package project

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
)

func (p *Project) Validate() {
	defer func() {
		if r := recover(); r != nil {
			p.concern = append(p.concern, constant.PanicOccurred)
		}
	}()
	versionSemantic := semver.New(p.Version)
	runtimeSemantic := semver.New(p.runtimeVersion)

	if versionSemantic.LessThan(*runtimeSemantic) {
		p.concern = append(p.concern, constant.RuntimeOld)
	}

	if runtimeSemantic.LessThan(*versionSemantic) {
		p.concern = append(p.concern, constant.RuntimeNewer)
	}
}
