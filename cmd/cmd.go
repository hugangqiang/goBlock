package cmd

import (
	"block.io/core"
	)

func init()  {
	bc := core.NewBlockChain()
	bc.SendData("第一个区块第一个区块第一个区块第一个区块第一个区块第一个区块")
	bc.SendData("第二个区块第二个区块第二个区块第二个区块第二个区块第二个区块")
	bc.Print()
}
