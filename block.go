package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"log"
	"os"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	return result.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	return &block
}

func (b *Block) Print() {
	fmt.Println("*******************")
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("Prev: %x\n", b.PrevBlockHash)
	fmt.Printf("Hash: %x\n", b.Hash)
	fmt.Println("*******************")
}
