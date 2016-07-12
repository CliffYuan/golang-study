package main
import (
	"encoding/json"
	"fmt"
	"os"
)

type Response struct {
	Page int `json:"page"`
	Fruits []string
}

func main() {
	//有目标对象
	res:=&Response{
		Page:23,
		Fruits:[]string{"apple","pear"}}
	jsonStr,_:=json.Marshal(res)
	fmt.Println(string(jsonStr))//注意string()函数

	fmt.Println("--------------------")
	//
	byt:=[]byte(`{"num":13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err:=json.Unmarshal(byt,&dat);err!=nil{
		panic(err)
	}
	fmt.Println(dat)
	//获取
	strss:=dat["strs"].([]interface{})
	str1:=strss[0].(string)
	fmt.Println(str1)


	fmt.Println("--------------------")
	//
	strsss:=`{"page":1,"Fruits":["apple","peach"]}`
	resp:=Response{}
	json.Unmarshal([]byte(strsss),&resp)//转化为byte
	fmt.Println(resp)
	fmt.Println(resp.Fruits[0])

	fmt.Println("--------------------")
	//
	enc:=json.NewEncoder(os.Stdout)
	d:=map[string]int{"apple":5,"lettuce":7}
	enc.Encode(d)

}
