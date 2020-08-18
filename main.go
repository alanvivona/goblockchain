package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/alanvivona/blockchaingo/blockchain"
)

func main() {

	// create a new blockchain and add the 1st block "Genesis"
	bc := blockchain.Chain{}
	bc.Init([]byte("Genesis Block!"))

	// add 9 more blocks to the chain
	currentIndex := len(bc.Blocks)
	endIndex := 4
	logrus.Infof("Adding %d blocks to the chain", endIndex-currentIndex)
	for i := currentIndex + 1; i <= endIndex; i++ {
		message := fmt.Sprintf("Block number %d!", i)
		bc.AddBlock([]byte(message))
	}

	logrus.Info("Validating and printing the blockchain")
	fmt.Println()
	for _, block := range bc.Blocks {
		fmt.Printf("Block : '%s'\n", block.Data)
		fmt.Printf("\t Hash: %x\n", block.Hash)
		fmt.Printf("\t Link: %x\n", block.Link)
		fmt.Printf("\t Nonce: %v\n", block.Nonce)
		fmt.Printf("\t Valid: %v\n", blockchain.IsValid(block))
		fmt.Println()
	}

}
