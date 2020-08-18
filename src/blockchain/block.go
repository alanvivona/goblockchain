package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Block struct {
	Transactions []*Transaction //	this block's transactions
	Hash         []byte         //	this block's hash
	Link         []byte         //	the hash of the last block in the chain. this is the key part that links the blocks together
	Nonce        int64          //	the nonce used to sing the block. useful for verification
}

func (b *Block) Build(transactions []*Transaction, link []byte) {
	b.Transactions = transactions
	b.Link = link
	pow := &ProofOfWork{Block: b, Target: getProofOfWorkTarget()}
	b.Nonce, b.Hash = pow.Run()
}

func (b *Block) HashTransactions() []byte {
	hashes := [][]byte{}
	for _, t := range b.Transactions {
		hashes = append(hashes, t.ID)
	}
	transactionHash := sha256.Sum256(bytes.Join(hashes, []byte{}))
	return transactionHash[:]
}

func (b *Block) Serialize() ([]byte, error) {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(b)
	if err != nil {
		logrus.Error("Failed to serialize block: ", b, err)
	}
	return buffer.Bytes(), err
}

func (b *Block) Deserialize(data []byte) error {
	buffer := bytes.Buffer{}
	buffer.Write(data)
	decoder := gob.NewDecoder(&buffer)
	err := decoder.Decode(&b)
	if err != nil {
		logrus.Error("Failed to deserialize data into block: ", data, err)
	}
	return err
}

func (b *Block) Print() {
	fmt.Printf("\t Transactions:\t%s\n", len(b.Transactions))
	fmt.Printf("\t Hash:\t%x\n", b.Hash)
	fmt.Printf("\t Link:\t%x\n", b.Link)
	fmt.Printf("\t Nonce:\t%d\n", b.Nonce)
}
