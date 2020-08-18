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
	genesisBlock := makeGenesisBlock()
	c.storage.Init(persistance.DefaultPath, genesisBlock, genesisBlock.Hash)
	return nil
}

func makeGenesisBlock() *Block {
	newBlock := &Block{}
	emptyLink := []byte{}
	newBlock.Build([]byte("Genesis Block!"), emptyLink)
	return newBlock
}

func (c *Chain) AddBlock(data []byte) error {

	lastHash, err := c.storage.GetLastHash()
	if err != nil {
		logrus.Error("Failed to get last hash from the storage: ", err)
		return err
	}
	newBlock := &Block{}
	newBlock.Build(data, lastHash)
	err = c.storage.SaveBlock(newBlock.Hash, newBlock)
	if err != nil {
		logrus.Error("Failed to save block into the storage: ", newBlock, err)
		return err
	}
	c.LastHash = newBlock.Hash
	return nil
}

func (c *Chain) IterateHashSort() error {
	logrus.Info("Iterating over the blockchain by hash value order...")
	block := &Block{}
	// This prefix filters out the metadata keys. All the hashes from the blocks begin with 0x00
	prefix := []byte{00}
	err := c.storage.Iterate(prefix, block, func(data []byte) error {
		if err := block.Deserialize(data); err != nil {
			return err
		}
		logrus.Info("--------")
		block.Print()
		return nil
	})
	if err != nil {
		logrus.Error("Error occurred during blockchain iteration: ", err)
		return err
	}
	logrus.Info("--------")
	return nil
}

func (c *Chain) IterateLink() error {
	logrus.Info("Iterating over the blockchain by link order...")
	currentHash := c.LastHash
	for currentHash != nil && len(currentHash) > 0 {
		data, err := c.storage.Get(currentHash)
		if err != nil {
			return err
		}
		block := &Block{}
		if err = block.Deserialize(data); err != nil {
			return err
		}
		logrus.Info("--------")
		block.Print()
		currentHash = block.Link
	}
	logrus.Info("--------")
	return nil
}
