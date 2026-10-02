package dependency

import (
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/license"
	"github.com/funtimecoding/soil/pkg/system"
	"slices"
)

func (d *Dependency) Validate() {
	defer func() {
		if r := recover(); r != nil {
			d.concern = append(d.concern, constant.PanicOccurred)
		}
	}()
	d.Files = license.FindFiles(d.Directory)

	if len(d.Files) == 0 {
		d.Family = constant.LicenseUnknown
		d.concern = append(d.concern, constant.LicenseMissing)

		return
	}

	for _, f := range d.Files {
		for _, i := range license.Scan(system.ReadFile(d.Directory, f)) {
			if !slices.Contains(d.Identifiers, i) {
				d.Identifiers = append(d.Identifiers, i)
			}
		}
	}

	d.Family = license.Family(d.Identifiers)

	switch d.Family {
	case constant.LicenseUnknown:
		d.concern = append(d.concern, constant.LicenseUnknown)
	case constant.WeakCopyleft:
		d.concern = append(d.concern, constant.WeakCopyleft)
	case constant.Copyleft:
		d.concern = append(d.concern, constant.Copyleft)
	}
}
