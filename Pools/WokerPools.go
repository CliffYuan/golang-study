package main
import (
	"fmt"
	"time"
)

func main() {
	jobs:=make(chan int,100)
	result:=make(chan int,100)

	for w := 1; w <= 3; w++ {
		go worker(w,jobs,result)
	}

	for j := 1; j <= 9; j++ {
		jobs<-j
	}

	close(jobs)

	fmt.Println("----------------")
	for a := 1; a <= 9; a++ {
		fmt.Println(<-result)
	}


}


func worker(id int, jobs <-chan int,result chan <-int)  {
	fmt.Println("worker ",id," start...")
	for j := range jobs {
		fmt.Println("worker ",id," processing job",j)
		time.Sleep(time.Second)
		result<-j
	}
}
