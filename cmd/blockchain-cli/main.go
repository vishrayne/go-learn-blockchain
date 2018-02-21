package main

import (
	"flag"
	"fmt"
	"os"

	blockchain "github.com/vishrayne/go-learn-blockchain"
)

// CLI will be our command line abstraction
type CLI struct {
	bc *blockchain.Blockchain
}

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.Close()

	cli := CLI{bc}
	cli.run()
}

func (cli *CLI) run() {
	cli.validateArgs()

	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "<BLOCK_DATA>")
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	var err error
	switch os.Args[1] {
	case "help":
		err = helpCmd.Parse(os.Args[2:])
	case "addblock":
		err = addBlockCmd.Parse(os.Args[2:])
	case "printchain":
		err = printChainCmd.Parse(os.Args[2:])
	}

	if err != nil {
		fmt.Printf("Error occured while executing command => %v", err)
		cli.printUsage()
		os.Exit(1)
	}

	if helpCmd.Parsed() {
		cli.printUsage()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}

		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\taddblock -data <BLOCK_DATA> - add the given block to the blockchain")
	fmt.Println("\tprintchain - print all the blocks of the chain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	if len(data) <= 0 {
		fmt.Println("Empty block data, skipping!")
		os.Exit(1)
	}

	cli.bc.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

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
