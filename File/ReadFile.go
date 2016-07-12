package main
import (
	"io/ioutil"
	"fmt"
	"os"
	"io"
	"bufio"
)


func checkErr(e error)  {
	if e!=nil{
		panic(e)
	}
}

func main() {
	//readfile
	data,err:=ioutil.ReadFile("/tmp/godata")
	checkErr(err)
	fmt.Println(string(data))

	fmt.Println("----------")

	//open
	f,err:=os.Open("/tmp/godata")
	defer f.Close()
	checkErr(err)
	b1:=make([]byte,5)
	n1,err:=f.Read(b1)
	checkErr(err)
	fmt.Printf("%d bytes:%s\n",n1,string(b1))

	fmt.Println("----------")
	//seek
	o2,err:=f.Seek(6,0)
	checkErr(err)
	b2:=make([]byte,2)
	n2,err:=f.Read(b2)
	checkErr(err)
	fmt.Printf("%d bytes @%d:%s\n",n2,o2,string(b2))

	fmt.Println("----------")
	//
	o3,err:=f.Seek(6,0)
	checkErr(err)
	b3:=make([]byte,2)
	n3,err:=io.ReadAtLeast(f,b3,2)
	checkErr(err)
	fmt.Printf("%d bytes @%d:%s\n",n3,o3,string(b3))


	fmt.Println("----------")
	_,err1:=f.Seek(0,0)
	checkErr(err1)
	r4:=bufio.NewReader(f)
	b4,err:=r4.Peek(5)
	checkErr(err)
	fmt.Printf("5 bytes:%s\n",string(b4))




}
