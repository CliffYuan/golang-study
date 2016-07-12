package main
import (
	"time"
	"fmt"
)

func main() {
	ticker1:=time.NewTicker(time.Millisecond*500)

	go func() {
		for t := range ticker1.C {
			fmt.Println("ticker at",t)
		}
	}()


	time.Sleep(time.Millisecond*1600)
	ticker1.Stop()
	fmt.Println("ticker stopped...",time.Now())
}
