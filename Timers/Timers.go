package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	timer1:=time.NewTimer(time.Second*2)
	<-timer1.C
	fmt.Println(time.Now(),"timer 1 expired")


	//
	timer2:=time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println(time.Now(),"timer 2 expired")
	}()
	stop2:=timer2.Stop()
	if stop2{
		fmt.Println(time.Now(),"time2 is stopped")
	}


}
