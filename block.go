package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

// Block representation
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// IsValidBlock will validate a block
func (b *Block) IsValidBlock() bool {
	pow := newPow(b)
	return pow.validate()
}

func (b *Block) serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		fmt.Printf("Unable to serialize block ==> %v", err)
		return nil
	}

	return result.Bytes()
}

func deserializeBlock(data []byte) *Block {
	var block *Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		fmt.Printf("Unable to de-serialize data ==> %v", err)
		return nil
	}

	return block
}

func newBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	block.setHash()
	return block
}

func newGenesisBlock() *Block {
	return newBlock("Genesis block", []byte{})
}

func (b *Block) setHash() {
	pow := newPow(b)
	nonce, hash := pow.run()
	b.Nonce = nonce
	b.Hash = hash[:]
}
