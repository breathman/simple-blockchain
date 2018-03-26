package main

func main() {
	bc := NewBlockchain()
	bc.AddBlock([]byte("Hello world"))
	bc.Print()
}
