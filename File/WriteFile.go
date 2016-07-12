package main
import (
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
)

func checkErr(e error)  {
	if e!=nil{
		panic(e)
	}
}

func main() {

	d1:=[]byte("hello\ngo\n")
	err:=ioutil.WriteFile("/tmp/godataw",d1,0644)
	checkErr(err)

	f,err:=os.Create("/tmp/godataw2")
	checkErr(err)
	defer f.Close()

	d2:=[]byte{115, 111, 109, 101, 10}
	n2,err:=f.Write(d2)
	checkErr(err)
	fmt.Printf("wrote %d bytes\n",n2)

	n3,err:=f.WriteString("writes\n")
	checkErr(err)
	fmt.Printf("wrote %d bytes\n",n3)
	f.Sync()

	//buffer
	w:=bufio.NewWriter(f)
	n4,err:=w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n",n4)
	w.Flush()
}
