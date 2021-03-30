package main

import (
	"fmt"
	"time"
)

func main() {
	//testTick()
	//	testNewTicker()
	//testAfter()
	testMyTimer()
}

func testAfter() {
	tchan := time.After(time.Second * 3)
	fmt.Println(time.Now().String(), "tchan=", <-tchan)
}

func testNewTicker() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		case <-done:
			return
		}
	}
}

func testTick() {
	c := time.Tick(2 * time.Second)
	for next := range c {
		fmt.Printf("%v \n", next)
	}
}

type MyTimer struct {
	timeDelay time.Duration
	c         chan time.Time
}

func NewTimer(t time.Duration) *MyTimer {
	return &MyTimer{t, make(chan time.Time)}
}

func (t *MyTimer) Tick() {
	go func() {
		for {
			time.Sleep(t.timeDelay)
			t.c <- time.Now()
		}
	}()
}

func testMyTimer() {
	timer := NewTimer(2 * time.Second)
	timer.Tick()
	for i := range timer.c {
		fmt.Println(i)
	}
}
