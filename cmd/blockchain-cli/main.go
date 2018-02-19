package main

import (
	"fmt"

	blockchain "github.com/vishrayne/go-learn-blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	bc.AddBlock("First transaction")
	// bc.AddBlock("Second transaction")
	// bc.AddBlock("Third transaction")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev block hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("PoW: %t\n\n", block.IsValidBlock())
	}
}
