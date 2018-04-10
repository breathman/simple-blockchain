package main

import (
	"os"
	"fmt"
	"flag"
	"log"
)

type CLI struct {
	bc *Blockchain
}

func (cli *CLI) Run() {
	var err error
	cli.ValidateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data","","Block data")

	switch os.Args[1] {
	case "addblock": {
		err = addBlockCmd.Parse(os.Args[2:])
	}
	case "printchain": {
		err = printChainCmd.Parse(os.Args[2:])
	}
	default:
		cli.PrintUsages()
		os.Exit(1)
	}
	if err!=nil {
		log.Fatal(err)
		cli.PrintUsages()
		os.Exit(2)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) ValidateArgs() {
	if len(os.Args) < 2 {
		cli.PrintUsages()
		os.Exit(1)
	}
}

func (cli *CLI) PrintUsages() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock([]byte(data))
}

func (cli CLI) printChain() {
	cli.bc.Print()
}