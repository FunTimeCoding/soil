package face

type CommandRecorder interface {
	BeginCommand(name string)
	RecordCommand(name string)
}
