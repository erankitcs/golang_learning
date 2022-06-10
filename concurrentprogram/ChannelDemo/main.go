package main

import (
	"fmt"
	"sync"
)

func main() {
	println("Channel Demo..")
	wg := &sync.WaitGroup{}
	//ch := make(chan int)
	//Buffered channel
	ch := make(chan int, 2)
	wg.Add(2)
	/* go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("Channel Reader..")
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 2
		ch <- 5
		wg.Done()
	}(ch, wg) */

	//Different channel type

	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println("Channel Reader Recieve only..")
		//fmt.Println(<-ch)
		//fmt.Println(<-ch)
		//Better way to read from channel.
		//if msg, ok := <-ch; ok {
		//	fmt.Println(msg)
		//}

		for i := range ch {
			fmt.Println(i)
		}

		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		//Channel push only
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//only normal channel or push only channel can close the channel.
		//close(ch)
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
