package blockchain

// Blockchain representation
type Blockchain struct {
	Blocks []*block
}

// NewBlockchain inits a new block chain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*block{newGenesisBlock()}}
}

// AddBlock adds a new block by string
func (bc *Blockchain) AddBlock(data string) {
	newBlock := newBlock(data, bc.getLastBlock().Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) getLastBlock() *block {
	return bc.Blocks[len(bc.Blocks)-1]
}
