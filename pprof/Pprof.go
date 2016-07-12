package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
	_ "net/http/pprof"
)

func main() {
	flag.Parse()

	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060",nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go work(&wg)
	}

	wg.Wait()

	time.Sleep(3*time.Second)

}

func work(wg *sync.WaitGroup)  {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1e10; i++ {
		time.Sleep(time.Millisecond*100)
		counter++
	}

	wg.Done()
}
