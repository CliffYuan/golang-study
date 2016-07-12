package main
import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h:=sha1.New()
	s:="sha1 this string"

	h.Write([]byte(s))

	bs:=h.Sum(nil)

	fmt.Printf("%x\n",bs)

}
