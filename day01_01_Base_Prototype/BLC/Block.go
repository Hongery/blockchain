package BLC

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
)

//1 创建Block结构体
type Block struct{
	Height         int64  //高度Height：其实就是区块的编号，第一个区块叫创世区块，高度为0
	PrevBlockHash  []byte  //上一个区块的哈希值ProvHash：
	Data 		   []byte  //交易数据Data：目前先设计为[]byte,后期是Transaction
	TimeStamp      int64	//时间戳TimeStamp：
	Hash            []byte	//哈希值Hash：32个的字节，64个16进制数
}

//2创建新的区
func NewBlock(data string,prevBlockHash []byte,height int64) *Block{
	//创建区块
	block:=&Block{height,prevBlockHash,[]byte(data),time.Now().Unix(),nil}
	//设置哈希值
	block.SetHash()
	return block
}


//3 设置区块的hash
func (block *Block)SetHash(){
	heightBytes :=IntToHex(block.Height)   //1.将高度转为字节数组
	timeString :=strconv.FormatInt(block.TimeStamp,2) //2.时间戳转为字节数组,转为二进制的字符串
	timeBytes :=[]byte(timeString)
	//3.拼接所有的属性
	blockBytes :=bytes.Join([][]byte{
		heightBytes,
		block.PrevBlockHash,
		block.Data,
		timeBytes},
		[]byte{})
	//4,生成hash值
	hash :=sha256.Sum256(blockBytes)  //数组长度32为
	block.Hash=hash[:]
}

//4.创建创世区块
func CreateGenesisBlock(data string) *Block{
	//因此make（[] int，0，10）分配长度为0的切片,容量10。 func make(t Type, size ...IntegerType) Type
	return  NewBlock(data,make([] byte,32,32),0)

}