package main

import (
	"fmt"

	blockchain "github.com/vishrayne/go-learn-blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.Close()

	bc.AddBlock("First transaction")
	// bc.AddBlock("Second transaction")
	// bc.AddBlock("Third transaction")

	bci := bc.Iterator()

	for {
		block := bci.Next()
		fmt.Printf("Prev block hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("PoW: %t\n\n", block.IsValidBlock())

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
