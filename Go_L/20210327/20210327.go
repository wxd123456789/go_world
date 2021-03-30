package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//mySet := New()
	//mySet.Add("a")
	//mySet.Add("b")
	//mySet.Add(1)
	//print(mySet)
	httpClientLongCon()
}

func httpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "hhh")
	})
	_ = http.ListenAndServe("127.0.0.0.1:8080", nil)
}

type Set struct {
	e map[interface{}]bool
}

//对外暴露的构造函数
func New() *Set {
	return &Set{e: make(map[interface{}]bool)}
}

func (set *Set) Add(element interface{}) bool {
	if !set.e[element] {
		set.e[element] = true
		return true
	}
	return false
}

func (set *Set) Remove(element interface{}) {
	delete(set.e, element)
}

func (set *Set) Clear() {
	set.e = make(map[interface{}]bool)
}

func (set *Set) Contains(element interface{}) bool {
	return set.e[element]
}

func (set *Set) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	for k := range set.e {
		buf.WriteString(fmt.Sprintf("%v,", k))
	}
	buf.WriteString("}")
	return buf.String()
}

func print(o ...interface{}) {
	fmt.Println(o)
}

/////长连接：客户端发送RESTFUL请求，需要监测某一资源变化情况，服务端提供watch机制，在资源有变化时通知client端。
func httpClientLongCon() {
	req, err := http.NewRequest("GET", "https://www.baidu.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	httpClient := &http.Client{}
	ret, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 4096)
	for {
		n, err := ret.Body.Read(buf)
		if n == 0 && err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
