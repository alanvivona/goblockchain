package blockchain

type Block struct {
	Data  []byte //	this block's data
	Hash  []byte //	this block's hash
	Link  []byte //	the hash of the last block in the chain. this is the key part that links the blocks together
	Nonce int64  //	the nonce used to sing the block. useful for verification
}

func (b *Block) Build(data []byte, link []byte) {
	b.Data = data
	b.Link = link
	pow := &ProofOfWork{Block: b, Target: getProofOfWorkTarget()}
	b.Nonce, b.Hash = pow.Run()
}
