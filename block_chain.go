package blockchain

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// Blockchain representation
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

// Iterator representation
type Iterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Iterator will iterate over the current blockchain
func (bc *Blockchain) Iterator() *Iterator {
	return &Iterator{bc.tip, bc.db}
}

// Close will perform necessary cleanup
func (bc *Blockchain) Close() {
	bc.db.Close()
}

// Next will move on to the next block in the chain
func (i *Iterator) Next() *Block {
	var block *Block
	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get([]byte(i.currentHash))
		block = deserializeBlock(encodedBlock)
		return nil
	})

	if err != nil {
		fmt.Printf("Unable to iterate to next block => %v", err)
		return nil
	}

	i.currentHash = block.PrevBlockHash

	return block
}

// NewBlockchain inits a new block chain
func NewBlockchain() *Blockchain {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	var tip []byte

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := newGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}

			err = b.Put(genesis.Hash, genesis.serialize())
			if err != nil {
				return err
			}

			err = b.Put([]byte("1"), genesis.Hash)
			if err != nil {
				return err
			}

			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("1"))
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Unable to update database => %v", err)
		return nil
	}

	return &Blockchain{tip, db}
}

// AddBlock adds a new block by string
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("1"))
		return nil
	})

	if err != nil {
		fmt.Printf("Unable to add a new block => %v", err)
		return
	}

	newBlock := newBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.serialize())
		if err != nil {
			return err
		}

		err = b.Put([]byte("1"), newBlock.Hash)
		if err != nil {
			return err
		}

		bc.tip = newBlock.Hash
		return nil
	})

	if err != nil {
		fmt.Printf("Unable to add a new block => %v", err)
		return
	}
}
