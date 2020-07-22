package block

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Data []byte //	this block's data
	Hash []byte //	this block's hash
	Link []byte //	the hash of the last block in the chain. this is the key part that links the blocks together
}

func (b *Block) DeriveHash() []byte {
	blockInfo := bytes.Join([][]byte{b.Data, b.Link}, []byte{})
	hash := sha256.Sum256(blockInfo)
	return hash[:]
}

func New(data []byte, link []byte) *Block {
	block := &Block{
		Data: data,
		Link: link,
	}
	block.Hash = block.DeriveHash()
	return block
}
