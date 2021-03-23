package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

var a int

func main() {
	fmt.Println("calm down i can")
	fmt.Println(MD5V([]byte("123456")))
	go func() {
		a++
	}()
	time.Sleep(time.Second)
	fmt.Println(a)
}

func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
