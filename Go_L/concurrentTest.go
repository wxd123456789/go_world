package main

import (
	. "fmt"
)

func run() {
	for i := 0; i < 5; i++ {
		Printf("hello " + string(i))
	}
}
func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func fibo(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x = y
		y = x + y
	}
	close(c)
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit:
			Println("quit")
			return
		}
	}
}


func main() {
	/*go run()
	println("end")
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c
	Println(x, y, x+y)

	e := make(chan int, 2)
	e <- 1
	e <- 2
	Println(<-e)
	Println(<-e)

	c = make(chan int, 10)
	go fibo(cap(c), c)
	for i := range c {
		Println(i)
	}*/
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
