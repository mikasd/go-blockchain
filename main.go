package main

import (
	"fmt"
	"github.com/mikasd/go-blockchain/blockchain"
)


func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("second block")
	chain.AddBlock("third block")
	chain.AddBlock("fourth block")

	for _, block := range chain.Blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Current Hash: %x\n", block.Hash)
	}
}
