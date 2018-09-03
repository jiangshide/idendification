package core

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Block struct {
	TimeStamp    int64
	Data         [] byte
	PreBlockHash []byte
	Hash         []byte
	Nonce        int
}

func(b *Block) Serialize()[]byte{
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil{
		log.Panic(err)
	}
	return result.Bytes()
}

func NewBlock(data string,preBlockHash []byte)*Block{
	block := &Block{time.Now().Unix(),[]byte(data),preBlockHash,[]byte{},1}
	pow := NewProofOfWork(block)
	nonce,hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	//block.SetHash()
	return block
}

func (b *Block) SetHash(){
	timeStamp := []byte(strconv.FormatInt(b.TimeStamp,10))
	headers := bytes.Join([][]byte{b.PreBlockHash,b.Data,timeStamp},[]byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewGensisBlock() *Block{
	return NewBlock("Gensis Block",[]byte{})
}

func DeserializeBlock(d []byte) *Block{
	var block Block
	decoder := gob.NewDecoder(bytes.NewBuffer(d))
	err := decoder.Decode(&block)
	if err != nil{
		log.Panic(err)
	}
	return &block
}