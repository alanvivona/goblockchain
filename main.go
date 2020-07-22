package main

import (
	"fmt"

	"github.com/alanvivona/blockchaingo/chain"
	"github.com/alanvivona/blockchaingo/ui"
)

func main() {

	// create a new blockchain and add the 1st block "Genesis"
	blockchain := chain.New()

	// add 9 more blocks to the chain
	for i := 2; i < 11; i++ {
		message := fmt.Sprintf("Block number %d!", i)
		blockchain.AddBlock([]byte(message))
	}

	// print the whole blockchain
	for _, block := range blockchain.Blocks {
		ui.PrintBlock(block)
	}

}
