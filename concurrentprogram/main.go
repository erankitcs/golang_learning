package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/erankitcs/golang_learning/concurrentprogram/books"
)

var cache = map[int]books.Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	println("Welcome to Go concurrent programming.")
	//Simple without go routine.
	/* for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		if b, ok := QueryCache(id); ok {
			fmt.Println("from cache")
			println(b.String())
			continue
		}

		if b, ok := QueryDatabase(id); ok {
			fmt.Println("from Database")
			cache[id] = b
			println(b.String())
			continue
		}
		fmt.Printf("Book not found id: '%v'", id)
	} */

	//With Go Routine
	/* for i := 0; i < 10; i++ {
		println(i)
		id := rnd.Intn(10) + 1
		go func(id int) {
			if b, ok := QueryCache(id); ok {
				fmt.Println("from cache")
				println(b.String())
			}
		}(id)

		go func(id int) {
			if b, ok := QueryDatabase(id); ok {
				fmt.Println("from Database")
				// cache[id] = b
				println(b.String())
			}
		}(id)
		time.Sleep(150 * time.Millisecond)

	}

	time.Sleep(2 * time.Second) */

	///Go routine with wait group
	/* wg := &sync.WaitGroup{}
	//m := &sync.Mutex{}
	m := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		println(id)
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := QueryCache(id, m); ok {
				fmt.Println("from cache")
				println(b.String())
			}
			wg.Done()
		}(id, wg, m)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := QueryDatabase(id); ok {
				fmt.Println("from Database")
				m.Lock()
				cache[id] = b
				m.Unlock()
				println(b.String())
			}
			wg.Done()
		}(id, wg, m)

		time.Sleep(150 * time.Millisecond)
	} */

	///goroutine and Mutex and Channels

	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan books.Book)
	dbCh := make(chan books.Book)

	for i := 0; i < 10; i++ {
		wg.Add(2)
		id := rnd.Intn(10) + 1
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- books.Book) {
			if b, ok := QueryCache(id, m); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- books.Book) {
			if b, ok := QueryDatabase(id); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbCh)

		go func(cacheCh, dbCh <-chan books.Book) {
			/// This i blocking select. Means if one channel is not ready then it will fail. Use
			// Default for non blocking select statement.
			select {
			case b := <-cacheCh:
				fmt.Println("From Cache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("From Database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func QueryCache(id int, m *sync.RWMutex) (books.Book, bool) {
	//m.Lock()
	m.RLock()
	b, ok := cache[id]
	//m.Unlock()
	m.RUnlock()
	return b, ok
}

func QueryDatabase(id int) (books.Book, bool) {
	time.Sleep(100 * time.Microsecond)

	for _, b := range books.MyBook {
		if b.ID == id {
			return b, true
		}
	}

	return books.Book{}, false
}
