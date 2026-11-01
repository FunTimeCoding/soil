package pipeline_view

func New(
	instance string,
	project int64,
	pipeline int64,
) *View {
	return &View{Instance: instance, Project: project, Pipeline: pipeline}
}
