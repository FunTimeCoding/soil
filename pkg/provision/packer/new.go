package packer

import (
	"github.com/funtimecoding/soil/pkg/provision/packer/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func New(workDirectory string) *Client {
	web := join.Absolute(
		workDirectory,
		constant.PackerDirectory,
		constant.WebDirectory,
	)
	system.MakeDirectory(web)

	return &Client{
		workDirectory:      workDirectory,
		packerWebDirectory: web,
		packerOutputDirectory: join.Absolute(
			workDirectory,
			constant.PackerDirectory,
			constant.OutputDirectory,
		),
	}
}
