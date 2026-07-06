package scan

import "github.com/funtimecoding/soil/pkg/lint/concern"

type Service struct {
	Name              string
	Repo              string
	ModelContext      bool
	Server            bool
	Web               bool
	Store             bool
	Generated         bool
	Convert           bool
	Client            bool
	Types             bool
	Model             bool
	ConstantDirectory bool
	ConstantFile      bool
	Worker            bool
	IntegrationTests  bool
	Option            bool
	Run               bool
	Concerns          []*concern.Concern
}
