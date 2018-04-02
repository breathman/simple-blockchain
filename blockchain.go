package main

import (
	"time"
)

type Blockchain struct {
	blocks []*Block
}

func NewBlockchain() (bc *Blockchain) {
	bc = &Blockchain{}
	bc.genesis()
	return
}

func (bc *Blockchain) genesis() {
	bc.blocks = append(bc.blocks, NewBlock([]byte("Genesis block"), []byte{}))
}

func NewBlock(data []byte, prev []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prev, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.nonce = nonce
	block.hash = hash
	return block
}

func (bc *Blockchain) AddBlock(data []byte) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) GetLastBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}

func (bc *Blockchain) Print() {
	for _, block := range bc.blocks {
		block.Print()
	}
}
