package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func testJson() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	/*	_s := "ABC"
		b := []byte(_s)
		fmt.Println(b)
		a := 10
		as := float32(a)
		fmt.Println(as)*/
}

func makeFile() {
	os.Mkdir("tmptest", 0777)
	fileName := "tmp.txt"
	_f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(_f, err)
		return
	}
	defer _f.Close()
	for i := 0; i < 10; i++ {
		_f.WriteString("test\n")
		_f.Write([]byte("test\n"))
	}
}

func readFile() {
	fileName := "tmp.txt"
	_f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(_f, err)
		return
	}
	defer _f.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := _f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func testStrOp() {
	// str operation
	fmt.Println(strings.Contains("seafood", "foo"))

	// *** to str
	str := make([]byte, 0, 100) //byte slice
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	// str to ***
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	fmt.Println(a, b)
}
func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	//makeFile()
	//readFile()
	testStrOp()
}
