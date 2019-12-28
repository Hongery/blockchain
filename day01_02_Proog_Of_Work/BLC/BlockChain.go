package BLC


//5创建区块连

type BlockChain struct {
	Blocks []*Block  //存储有序的区块
}

//step6：创建区块链，带有创世区块
func CreateBlockChainWithGenesisBlock(data string ) *BlockChain{
	//创建创世区块
	genesisBlock :=CreateGenesisBlock(data)
	//返回区块连对象
	return &BlockChain{[]*Block{genesisBlock}}
}

//7 添加一个新区快
func (bc *BlockChain) AddBlockToBlockChain(data string,height int64,prevHash []byte){
	//创建新区块
	newBlock :=NewBlock(data,prevHash,height)
	//添加到切片中
	bc.Blocks=append(bc.Blocks,newBlock)







}

