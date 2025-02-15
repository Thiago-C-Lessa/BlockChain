package main

import (
 	"Blockchain/blockchain"
	"fmt"
)



func main() {
	chain := block.InitBlockChain()

	chain.AddBlock("Primeiro")
	chain.AddBlock("Segundo")
	chain.AddBlock("Terceiro")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}