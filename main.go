package main

import "fmt"

func main() {
	blockchain := CreateBlockchain(3)

	// record transactions on the blockchain for rasa, dowi and uman
	blockchain.addBlock("raka", "uman", 0.001)
	blockchain.addBlock("uman", "dowi", 0.6)

	// check if the blockchain is valid: expecting true
	fmt.Println(blockchain.isValid())
}
