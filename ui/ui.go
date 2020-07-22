package ui

import (
	"encoding/hex"

	"github.com/alanvivona/blockchaingo/block"
	"github.com/sirupsen/logrus"
)

func HashBytesToString(hash []byte) string {
	encoded := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(encoded, hash)
	return string(encoded)
}

func PrintBlock(b *block.Block) {
	logrus.WithFields(logrus.Fields{
		"hash": HashBytesToString(b.Hash),
		"link": HashBytesToString(b.Link),
	}).Info(string(b.Data))
}
