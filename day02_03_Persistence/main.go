package main

import (

	"fmt"
	"mypublicchain/day02_03_Persistence/BLC"
)

func main() {

	//5.测试序列化和反序列化
	//block:=BLC.NewBlock("helloworld",make([]byte,32,32),0)
	//data:=block.Serilalize()
	//fmt.Println(block)
	//fmt.Println(data)
	//block2:=BLC.DeserializeBlock(data)
	//fmt.Println(block2)

	//7.测试创世区块存入数据库
	blockchain :=BLC.CreateBlockChainWithGenesisBlock("Genesis Block..")
	fmt.Println(blockchain)
	defer blockchain.DB.Close()
	//8.测试新添加的区块
	blockchain.AddBlockToBlockChain("Send 100RMB to wangergou")
	blockchain.AddBlockToBlockChain("Send 100RMB to lixiaohua")
	blockchain.AddBlockToBlockChain("Send 100RMB to rose")
	fmt.Println(blockchain)
	blockchain.PrintChains()

}