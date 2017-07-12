package main

import (
	"./helper"
	"fmt"
	"github.com/otium/queue"
	"time"
)

type A struct {
	t int
	s string
}

func main() {

	n := 20 // 并发处理数
	t1 := time.Now()
	fmt.Println("Hello, 世界")

	q := queue.NewQueue(func(val interface{}) {
		fmt.Println(val)
		time.Sleep(time.Second * 2)
	}, n)
	for i := 0; i < 20; i++ {
		a := A{t: i, s: helper.QuickRandom(i)}
		q.Push(a)
	}
	fmt.Println("Hello, 世界 2")
	q.Wait()
	fmt.Println("Hello, 世界3")
	t2 := time.Now()
	fmt.Println("消耗时间：", t2.Sub(t1), "秒")
}
