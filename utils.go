package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"slices"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func ReverseBytes(data []byte) {
	slices.Reverse(data)
}
