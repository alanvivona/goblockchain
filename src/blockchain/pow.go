/* Proof Of Work */
package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"

	"github.com/sirupsen/logrus"
)

/* In this first implementation the difficulty will remain constant
but in a real implementation we want this to be adjustable to fit our needs over time */
const Difficulty = int64(20)

type ProofOfWork struct {
	Block  *Block   // a block from the blockchain
	Target *big.Int // number that represents the requirements. it's derived form the difficulty
}

func getProofOfWorkTarget() *big.Int {
	target := big.NewInt(1)
	// left shift
	target.Lsh(target, uint(256-Difficulty))
	return target
}

func (pow *ProofOfWork) joinData(nonce int64) []byte {
	return bytes.Join(
		[][]byte{
			pow.Block.Link,
			pow.Block.HashTransactions(),
			big.NewInt(nonce).Bytes(),
			big.NewInt(Difficulty).Bytes(),
		},
		[]byte{},
	)
}

func (pow *ProofOfWork) Run() (int64, []byte) {
	var hash [32]byte
	var hashIntegerRep big.Int
	nonce := int64(0)

	logrus.WithFields(logrus.Fields{"block_transactions": len(pow.Block.Transactions)}).Info("Running Proof of Work...")
	for nonce < math.MaxInt64 {
		// the idea of this is to have, basically, an infinite loop
		// the nonce should be incremented every time
		// IMPLEMENT THIS IN ANOTHER WAY
		// something more gophery

		data := pow.joinData(nonce)
		hash = sha256.Sum256(data)
		hashIntegerRep.SetBytes(hash[:])
		fmt.Printf("\r%x", hash)

		// check if the hash is smaller than the target
		if hashIntegerRep.Cmp(pow.Target) == -1 {
			// block is signed
			break
		} else {
			// try next nonce
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

func IsValid(block *Block) bool {
	var hashIntegerRep big.Int
	pow := &ProofOfWork{Block: block, Target: getProofOfWorkTarget()}
	data := pow.joinData(block.Nonce)
	hash := sha256.Sum256(data)
	hashIntegerRep.SetBytes(hash[:])
	return hashIntegerRep.Cmp(pow.Target) == -1
}
