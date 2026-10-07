package recorded_commit

func New(
	branch string,
	message string,
) *Commit {
	return &Commit{Branch: branch, Message: message}
}
