package core

import (
	"log"
	"fmt"
)

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	genesisBlock := GenerateGenesisBlock()
	BlockChain := BlockChain{}
	BlockChain.ApendBlock(&genesisBlock)
	return &BlockChain
}

func (bc *BlockChain)SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.ApendBlock(&newBlock)
}

func (bc *BlockChain) ApendBlock(newBlock *Block){
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks) - 1 ]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("验证失败")
	}
}
func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Prev.Hash: %s\n", block.PrevBlockHash)
		fmt.Printf("Curr.Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %d\n\n", block.Timestamp)
	}
}
func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index -1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}