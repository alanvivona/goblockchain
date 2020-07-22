package chain

import (
	"github.com/alanvivona/blockchaingo/block"
)

type Chain struct {
	Blocks []*block.Block
}

func (c *Chain) AddBlock(data []byte) {
	lastBlock := c.Blocks[len(c.Blocks)-1]
	newBlock := block.New(data, lastBlock.Hash)
	c.Blocks = append(c.Blocks, newBlock)
}

func New() *Chain {
	return &Chain{
		Blocks: []*block.Block{
			block.New([]byte("Genesis Block"), []byte{}),
		},
	}
}
