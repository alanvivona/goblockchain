package main

import (
	"os"

	"github.com/alanvivona/blockchaingo/src/cli"

	"github.com/sirupsen/logrus"

	"github.com/alanvivona/blockchaingo/src/blockchain"
)

func main() {

	message, listLast, listAll, err := cli.Parse()
	if err != nil {
		logrus.Error("Failed to parse cli arguments: ", err)
		os.Exit(1)
	}

	bc := blockchain.Chain{}
	bc.Init()

	if message != nil && len(*message) > 0 {
		logrus.Info("Adding block to the chain: ", *message)
		err := bc.AddBlock([]byte(*message))
		if err != nil {
			logrus.Error("Failed to add block: ", err, message)
			return
		}
	}

	if listAll != nil && *listAll {
		cli.PrintLine()
		logrus.Info("= Blockchain =")
		err := bc.IterateLink(
			func(b *blockchain.Block) {
				b.Print()
				logrus.Info("--------")
			},
			func() { logrus.Info("--------") },
			func() { logrus.Info("End of blockchain") },
		)
		if err != nil {
			logrus.Error("Failed to iterate over blockchain: ", err)
			return
		}
		cli.PrintLine()
	} else if listLast != nil && *listLast {
		cli.PrintLine()
		logrus.Info("= Last Block =")
		lastBlock, err := bc.GetLastBlock()
		if err != nil {
			logrus.Error("Failed to get the last block: ", err)
			return
		}
		lastBlock.Print()
		cli.PrintLine()
	}

}
