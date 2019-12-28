package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

func main()  {
	/*
	1.安装数据库
		打开终端：go get "github.com/boltdb/bolt"
		此处需要稍微等待一下

	2.导入数据库的包
	 */

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//1.创建表
	//err = db.Update(func(tx *bolt.Tx) error{
	//	//1.创建MyBucket
	//	b,err := tx.CreateBucket([]byte("MyBycket"))
	//	if err != nil{
	//		return fmt.Errorf("create bucket:%s",err)
	//	}
	//	//2.向表中存储数据
	//	if b != nil{
	//		err := b.Put([] byte("l"),[] byte("send 100 BTC to 王二狗"))
	//		if err != nil{
	//			log.Panic("数据存储失败。。")
	//		}
	//	}
	//	return nil
	//})
	//
	//if err != nil{
	//	log.Panic(err)
	//}
	//读取数据
	err = db.View(func(tx *bolt.Tx) error {
		//获取bucket对象
		b := tx.Bucket([]byte("MyBycket"))
		if b != nil {
			//根据key查看数据
			data := b.Get([] byte("l"))//根据key获取对应的value值
			fmt.Println(data)
			fmt.Printf("%s\n", data)
			data2 := b.Get([] byte("ll"))//key不存在
			fmt.Println(data2)
			fmt.Printf("%s\n", data2) //[]，如果对应的key不存在，那么取出的是空。

		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	//读取数据库中的所有数据  读取key-value
	db.View(func(tx *bolt.Tx) error {
		b :=tx.Bucket([]byte("MyBycket"))
		c :=b.Cursor()
		for k,v:=c.First();k!=nil;k,v=c.Next(){
			fmt.Printf("key=%s,value=%s\n",k,v)
		}
		return nil
	})
}
