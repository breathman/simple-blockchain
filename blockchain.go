package main

import (
	"github.com/boltdb/bolt"
	"log"
	"os"
	"time"
)

const (
	dbFile       = "./bc.db"
	blocksBucket = "blocks"
)

type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

func NewBlockchain() *Blockchain {
	var tip []byte

	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))

		if bucket == nil {
			genesis := NewGenesisBlock()
			bucket, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}
			bucket.Put(genesis.Hash, genesis.Serialize())
			bucket.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = bucket.Get([]byte("l"))
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	return &Blockchain{tip, db}
}

func NewGenesisBlock() *Block {
	return NewBlock([]byte("Genesis block"), []byte{})
}

func NewBlock(data []byte, prev []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prev, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

func (bc *Blockchain) AddBlock(data []byte) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		lastHash = bucket.Get([]byte("l"))
		return nil
	})
	if err!=nil {
		log.Fatal(err)
		os.Exit(0)
	}

	newBlock := NewBlock(data, lastHash)


	err = bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put(newBlock.Hash, newBlock.Serialize())
		if err!=nil {
			return err
		}
		err = bucket.Put([]byte("l"), newBlock.Hash)
		if err!=nil {
			return err
		}
		return nil
	})
	if err!=nil {
		log.Fatal(err)
		os.Exit(0)
	}

}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.tip, bc.db}
}

func (bc *Blockchain) Print() {
	bci := bc.Iterator()

	for {
		block := bci.Next()
		block.Print()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}