package guard

func newHiddenGitCall() *GitCall {
	return &GitCall{hidden: true}
}
