package main

import (
	"fmt"
)

type Block struct {
	timestamp     int64
	data          []byte
	prevBlockHash []byte
	hash          []byte
	nonce         int
}

func (b *Block) Print() {
	fmt.Println("*******************")
	fmt.Printf("Data: %s\n", b.data)
	fmt.Printf("Prev: %x\n", b.prevBlockHash)
	fmt.Printf("Hash: %x\n", b.hash)
	fmt.Println("*******************")
}
