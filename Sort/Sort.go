package main
import (
	"sort"
	"fmt"
)

func main() {
	strs:=[]string{"d","b","a"}

	sort.Strings(strs)
	fmt.Println("strings:",strs)


	ints:=[]int{7,2,4}
	sort.Ints(ints)
	fmt.Println("ints:",ints)

	s:=sort.IntsAreSorted(ints)
	fmt.Println("sorted:",s)
}
