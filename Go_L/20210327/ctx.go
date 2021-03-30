package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//testValueCtx()
	//testCancelCtx()
	//testDeadlineCtx()
}

func testCancelCtx() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//每隔1s说一话，testCancelCtx函数在10s后执行cancel，那么speak检测到取消信号就会退出。
	go func(ctx context.Context) {
		for range time.Tick(time.Second) {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("speak")
			}
		}
	}(ctx)
	time.Sleep(10 * time.Second)
}

func testDeadlineCtx() {
	later, _ := time.ParseDuration("10s")
	deadline := time.Now().Add(later)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-time.After(20 * time.Second)://其他的case没有io的话，该io会阻塞20秒输入
			fmt.Println("stop")
		}
	}(ctx)
	time.Sleep(20 * time.Second)
}

func testTimeoutCtx() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case <-time.After(20 * time.Second):
			fmt.Println("stop monitor")
		}
	}(ctx)
	time.Sleep(20 * time.Second)
}

type key string

func testValueCtx() {
	ctx := context.WithValue(context.Background(), key("a"), "123")
	Get(ctx, "a")
	Get(ctx, "b")
}
func Get(ctx context.Context, k key) {
	if v, ok := ctx.Value(k).(string); ok {
		fmt.Println(v)
	}
}
