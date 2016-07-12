package main
import (
	"strings"
	"fmt"
)


func Maps(strs []string,f func(s string) string) []string {
	vsm:=make([]string,len(strs))
	for i, v := range strs {
		vsm[i]=f(v)
	}
	return vsm
}

func main() {
	strs:=[]string{"Apple","qq","Meizu"}
	s:=Maps(strs,strings.ToLower)
	for _, v := range s {
		fmt.Println(v)
	}
}
