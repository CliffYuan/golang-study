package main

import (
	"fmt"
	"time"
)

func main() {
	sliceA := []string{"A","B","C","D","E","F"}
	sliceB := sliceA[1:cap(sliceA)]
	fmt.Println(sliceB)

	fmt.Println("--------------append-----------------")

	//append返回一个新的切片了
	sliceC:=append(sliceA,"G","H")
	fmt.Println(sliceA)
	fmt.Println(sliceC)

	fmt.Println("--------------copy-----------------")

	//copy必须是相同类型
	//copy是覆盖同一个
	copyLen:=copy(sliceA,[]string{"G"})
	fmt.Println(sliceA)
	fmt.Println(copyLen)

	fmt.Println(time.Now())
}
