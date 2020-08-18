package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"errors"
)

/* REWORK THE MODULE TO PERSONAL STYLE/UNDERSTANDING */

const transactionMiningReward = 100

type transactionOutput struct {
	value     int
	publicKey string // first implementation: simple string user address
}

func (out *transactionOutput) CanBeUnlocked(data string) bool {
	// this is an over simplified impl
	return out.publicKey == data
}

type transactionInput struct {
	id  []byte
	out int
	sig string
}

func (in *transactionInput) CanUnlock(data string) bool {
	// this is an over simplified impl
	return in.sig == data
}

type Transaction struct {
	ID      []byte
	Inputs  []transactionInput
	Outputs []transactionOutput
}

func (t *Transaction) setID() error {
	var buffer bytes.Buffer

	err := gob.NewEncoder(&buffer).Encode(t)
	if err != nil {
		return err
	}

	hash := sha256.Sum256(buffer.Bytes())
	t.ID = hash[:]
	return nil
}

func (t *Transaction) IsCoinbase() bool {
	// coinbase transactions can only have one input and does not reference other transactions
	return len(t.Inputs) == 1 && len(t.Inputs[0].id) == 0 && t.Inputs[0].out == -1
}

func MakeCoinbaseTransaction(to, data string) (*Transaction, error) {
	if len(data) == 0 {
		return nil, errors.New("Empty data on coinbase transaction")
	}

	in := transactionInput{
		id:  []byte{}, // does not reference another transaction
		out: -1,       // does not reference another transaction
		sig: data,
	}
	out := transactionOutput{value: transactionMiningReward, publicKey: to}
	transaction := Transaction{
		ID:      nil,
		Inputs:  []transactionInput{in},
		Outputs: []transactionOutput{out},
	}
	if err := transaction.setID(); err != nil {
		return nil, err
	}
	return &transaction, nil
}
