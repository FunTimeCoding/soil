package symbol_rename

func New(
	oldName string,
	newName string,
	receiver string,
) *Rename {
	return &Rename{OldName: oldName, NewName: newName, Receiver: receiver}
}
