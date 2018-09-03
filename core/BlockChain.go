package core

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

const dbFile = "blockchain.db"
const blocksBucket  = "blocks"

type BlockChain struct {
	//Blocks []*Block
	tip []byte
	Db []bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	Db *bolt.DB
}

func (bc *BlockChain) AddBlock(data string){
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data,prevBlock.Hash)
	bc.Blocks = append(bc.Blocks,newBlock)
}

func NewBlockChain() *BlockChain{
	var tip []byte
	db,err := bolt.Open(dbFile,0600,nil)
	if err != nil{
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil{
			fmt.Println("No existing blockchain found. Creating a new one...")
			genesis := NewGensisBlock()
			b,err := tx.CreateBucket([]byte(blockBucket))
			if err != nil{
				log.Panic(err)
			}
			err = b.Put(genesis.Hash,genesis.Serialize())
			if err != nil{
				log.Panic(err)
			}
			err = b.Put([]byte("l"),genesis.Hash)
			if err != nil{
				log.Panic(err)
			}
			tip  = genesis.Hash
		}else{
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil{
		log.Panic(err)
	}
	bc := BlockChain{tip,db}
	return &bc
}

//func NewBlockChain()*BlockChain  {
//	return &BlockChain{[]*Block{NewGensisBlock()}}
//}
