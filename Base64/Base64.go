package main
import (
	"encoding/base64"
	"fmt"
)

func main() {
	data:="abc123!@#$%^^^^"
	en:=base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(en)

	fmt.Println("-----------")
	enDec,_:=base64.StdEncoding.DecodeString(en)
	fmt.Println(string(enDec))
	fmt.Println(enDec)

	fmt.Println("-----------")




}
