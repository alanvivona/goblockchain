package blockchain

import (
	"github.com/alanvivona/blockchaingo/src/persistance"
	"github.com/sirupsen/logrus"
)

type Chain struct {
	LastHash []byte
	storage  *persistance.Persistance
}

func (c *Chain) Init() error {
	logrus.WithFields(logrus.Fields{"difficulty": Difficulty}).Info("Initializing the blockchain...")
	c.storage = &persistance.Persistance{}
	c.storage.Init(persistance.DefaultPath, func() (persistance.Serializable, []byte) {
		genesisBlock := makeGenesisBlock()
		return genesisBlock, genesisBlock.Hash
	})
	return nil
}

func makeGenesisBlock() *Block {
	logrus.Info("Generating genesis block...")
	newBlock := &Block{}
	emptyLink := []byte{}
	newBlock.Build([]byte("Genesis Block!"), emptyLink)
	return newBlock
}

func (c *Chain) UpdateLastHash() error {
	logrus.Info("Fetchin last stored hash...")
	lastHash, err := c.storage.GetLastHash()
	if err != nil {
		logrus.Error("Failed to get last hash from the storage: ", err)
		return err
	}
	c.LastHash = lastHash
	return nil
}

func (c *Chain) AddBlock(data []byte) error {
	logrus.Info("Adding block to the blockchain...")
	if err := c.UpdateLastHash(); err != nil {
		return err
	}
	newBlock := &Block{}
	newBlock.Build(data, c.LastHash)
	err := c.storage.SaveBlock(newBlock.Hash, newBlock)
	if err != nil {
		logrus.Error("Failed to save block into the storage: ", newBlock, err)
		return err
	}
	c.LastHash = newBlock.Hash
	return nil
}

func (c *Chain) IterateLink(each func(b *Block), pre, post func()) error {
	logrus.Info("Iterating over the blockchain by link order...")
	c.UpdateLastHash()
	currentHash := c.LastHash
	pre()
	for currentHash != nil && len(currentHash) > 0 {
		data, err := c.storage.Get(currentHash)
		if err != nil {
			return err
		}
		block := &Block{}
		if err = block.Deserialize(data); err != nil {
			return err
		}
		each(block)
		currentHash = block.Link
	}
	post()
	return nil
}

func (c *Chain) GetLastBlock() (*Block, error) {
	c.UpdateLastHash()
	return c.GetBlock(c.LastHash)
}

func (c *Chain) GetBlock(hash []byte) (*Block, error) {
	logrus.Infof("Getting block %x form the storage...", hash)
	data, err := c.storage.Get(hash)
	if err != nil {
		return nil, err
	}
	block := &Block{}
	if err = block.Deserialize(data); err != nil {
		return nil, err
	}
	return block, nil
}
