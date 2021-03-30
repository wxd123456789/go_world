package main

import (
	"fmt"
	"time"
)

func main(){
	testExitG()
}

func run(done chan int){
	for{
		select {
		case <-done:
			fmt.Println("exit run")
			break
		default:
		}
		time.Sleep(time.Second)
		fmt.Println("do...")
	}
}

func testExitG(){
	done := make(chan int)
	go run(done)
	time.Sleep(2*time.Second)
	done<- 1
	fmt.Println("done")
}
