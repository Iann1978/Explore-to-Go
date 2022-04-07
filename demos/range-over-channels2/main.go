package main

import (
	"context"
	"fmt"
	"time"
)

func producer(ctx context.Context) <-chan int {

	queue := make(chan int, 10000)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				{
					fmt.Println("Done")
					close(queue)
					return
				}
			case queue <- n:
				n++
			}
		}
	}()

	return queue

}

func main() {

	ctx, cancel := context.WithCancel(context.TODO())
	queue := producer(ctx)
	// defer cancel()

	for elem := range queue {
		fmt.Println(elem)

		if elem == 2 {
			cancel()
		}
	}

	time.Sleep(time.Second)

}
