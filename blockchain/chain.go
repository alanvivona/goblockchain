package blockchain

import (
	"github.com/sirupsen/logrus"
)

type Chain struct {
	Blocks []*Block
}

func (c *Chain) Init(genesisBlockContent []byte) {
	logrus.WithFields(logrus.Fields{
		"content":    string(genesisBlockContent),
		"difficulty": Difficulty,
	}).Info("Initializing the blockchain")
	newBlock := &Block{}
	emptyLink := []byte{}
	newBlock.Build(genesisBlockContent, emptyLink)
	c.Blocks = []*Block{newBlock}
}

func (c *Chain) AddBlock(data []byte) {
	lastBlock := c.Blocks[len(c.Blocks)-1]
	newBlock := &Block{}
	newBlock.Build(data, lastBlock.Hash)
	c.Blocks = append(c.Blocks, newBlock)
}
