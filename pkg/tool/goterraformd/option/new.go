package option

import "github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"

func New() *Terraform {
	return &Terraform{Executable: constant.Command}
}
