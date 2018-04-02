package main

import "fmt"

func main() {

	bc := NewBlockchain()
	bc.AddBlock([]byte("Hello world"))

	for _, block := range bc.blocks {
		pow := NewProofOfWork(block)
		block.Print()
		fmt.Printf("Valid: %v \n\n", pow.Validate(block))
	}
}
