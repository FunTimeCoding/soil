package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/types/service"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/convert"
	"testing"
)

func TestConvertServicePublishesThePackage(t *testing.T) {
	v := convert.Service(
		service.New(
			"foxtrot.service",
			systemConstant.ServiceRunning,
			systemConstant.ServiceOriginVendor,
			"apt.example.org stable/main",
			"foxtrot",
			"1.2.3",
			true,
		),
	)
	assert.String(t, "foxtrot.service", v.Name)
	assert.String(t, "foxtrot", *v.Package)
	assert.String(t, "1.2.3", *v.Version)
}

func TestConvertServiceWithoutAPackagePublishesEmpty(t *testing.T) {
	v := convert.Service(
		service.New(
			"charlie.service",
			systemConstant.ServiceFailed,
			systemConstant.ServiceOriginLocal,
			"/etc/systemd/system",
			"",
			"",
			true,
		),
	)
	assert.String(t, "", *v.Package)
}
