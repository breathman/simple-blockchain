package main

import (
	"github.com/boltdb/bolt"
)

type BlockchainIterator struct {
	currentHash []byte
	db *bolt.DB
}

func (bci *BlockchainIterator) Next() *Block {
	var block *Block

	_ = bci.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		encodedBlock := bucket.Get(bci.currentHash)
		block = Deserialize(encodedBlock)

		return nil
	})

	bci.currentHash = block.PrevBlockHash

	return block
}