package blockchain

import (
	"time"
)

// Block representation
type block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func newBlock(data string, prevBlockHash []byte) *block {
	block := &block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	block.setHash()
	return block
}

func (b *block) IsValidBlock() bool {
	pow := newPow(b)
	return pow.validate()
}

func newGenesisBlock() *block {
	return newBlock("Genesis block", []byte{})
}

func (b *block) setHash() {
	pow := newPow(b)
	nonce, hash := pow.run()
	b.Nonce = nonce
	b.Hash = hash[:]
}
