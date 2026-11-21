package remediation_builder

func (b *Builder) LeaveOut(vulnerability string) {
	b.unimportant[vulnerability] = true
}
