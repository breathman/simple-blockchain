package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	hash := sha256.Sum256(bytes.Join([][]byte{timestamp, b.Data, b.PrevBlockHash}, []byte{}))
	b.Hash = hash[:]
}

func (b *Block) Print() {
	fmt.Println("*******************")
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("Prev: %x\n", b.PrevBlockHash)
	fmt.Printf("Block: %x\n", b.Hash)
	fmt.Println("*******************")
}
