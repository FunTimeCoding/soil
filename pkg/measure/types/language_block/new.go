package language_block

func New(
	open string,
	close string,
) *Block {
	return &Block{Open: open, Close: close}
}
