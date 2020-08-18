package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/alanvivona/blockchaingo/src/blockchain"
)

func main() {

	// create a new blockchain and add the 1st block "Genesis"
	bc := blockchain.Chain{}
	bc.Init()

	logrus.Info("Adding blocks to the chain")
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Block number %d!", i)
		bc.AddBlock([]byte(message))
	}

	bc.IterateLink()

}
