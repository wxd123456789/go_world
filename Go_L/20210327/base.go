package main

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	// orderIterMap()
	//a := []int{1, 2, 8,}
	//b := a[2:]
	//b[0] = 33
	//fmt.Println(a)
	//testPC()

	//a := []int{1, 2, 4}
	////删除队尾的元素
	//c := append(a[:2], nil...)
	//fmt.Printf("%v", c)
	//testSqQueue()

}

func testSelect() {
	a := make(chan int)
	select {
	default:
		fmt.Printf("default11")
	case <-a:
		fmt.Printf("recv sign")
	}
}

func testStrBuilder() {
	a := strings.Builder{}
	a.WriteString("123")
}

func testSqQueue() {
	sq := InitQueue()
	sq.EnQueue(1)
	sq.EnQueue(2)
}

const CAP = 5

type SqQueue struct {
	data  [CAP]int
	front int
	rear  int
}

func InitQueue() *SqQueue {
	return &SqQueue{front: 0, rear: 0}
}

func (s *SqQueue) Size() int {
	return s.rear - s.front + 1
}

func (s *SqQueue) EnQueue(d int) error {
	if (s.rear+1)%CAP == s.front {
		return errors.New("full")
	}
	s.data[s.rear] = d
	s.rear = (s.rear + 1) % CAP
	return nil
}
func (s *SqQueue) DeQueue() (int, error) {
	if s.rear == s.front {
		return 1, errors.New("empty")
	}
	e := s.data[s.front]
	s.data[s.front] = 0
	s.front = (s.front + 1) % CAP
	return e, nil
}

type queue struct {
	msg  []int
	lock sync.Mutex
}

func producer(q *queue, i int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.msg = append(q.msg, i)
}

func customer(q *queue) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for len(q.msg) > 0 {
		m := q.msg[0]
		fmt.Println(GetGID(), m)
		q.msg = q.msg[1:]
	}
}

func testPC() {
	q := queue{}
	for i := 0; i < 10; i++ {
		go producer(&q, i)
	}
	for i := 0; i < 10; i++ {
		go customer(&q)
	}
	time.Sleep(time.Second)
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

//chan
/*func producer(c chan int, i int) {
	c <- i
}

func customer(c chan int) {
	fmt.Println(<-c)
}

func testPC() {
	queue := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go producer(queue, i)
	}
	for i := 0; i < 10; i++ {
		go customer(queue)
	}
	time.Sleep(time.Second)
}
*/
func testEqual() {
	var a chan int
	var b chan int
	print(a == b)
	aA := [2]int{1, 2,}
	aB := [2]int{1, 2,}
	print(aA == aB)
	_p1 := p1{"1", 2, [2]int{1, 2},}
	_p2 := p1{"1", 2, [2]int{1, 2},}
	print(_p1 == _p2)
	_p3 := p2{"1", 2, [2]int{1, 2},}
	//强制类型转换
	print(p1(_p3) == _p2)
}

type p1 struct {
	f1 string
	f2 int
	f3 [2]int
}

type p2 struct {
	f1 string
	f2 int
	f3 [2]int
}

func orderIterMap() {
	m := make(map[string]string)
	m["a"] = "123"
	m["b"] = "456"
	keys := []string{
		"a", "b",
	}
	for i, _ := range keys {
		if value, ok := m[keys[i]]; ok {
			fmt.Println(value)
		} else {
			fmt.Printf("%s not exists", keys[i])
		}
	}
}
