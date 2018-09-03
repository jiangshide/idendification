package core

import (
	"bytes"
	"encoding/binary"
	"log"
	"crypto/sha256"
)

func IntToHex(num int64)[]byte{
	buff := new(bytes.Buffer)
	if err := binary.Write(buff,binary.BigEndian,num);err != nil{
		log.Panic(err)
	}
	return buff.Bytes()
}

func DataToHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}