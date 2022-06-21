package main

import (
	"fmt"
	"oneclick/src/Signature"
	"os"
)

func main() {
	const data = string (`["www.hitosea.com","www.json.cn","qq.com","baidu.com"]`)
	const datas = string (`["www.hitosea.com","www.json.cn","qq.com","baidu.com","zhihu.com","google.com"]`)
	ss := Signature.SignatureMethod{}.GeneralRequest("163793511","oneclick", "10")
	fmt.Println(ss)
	fmt.Fprintln(os.Stderr, ss)
	//c2a68de9054e4ba0bcbb6aa946771605
	//60  e1d68d4c756bb5d7d67b9ec4dec8f10f
	//360 775ccd59bcfc15d7a9ae812b984ff856
	//30 92475e9914fb16225bfd83f4de506725
	//test := make(map[string]string)
	//test["time"] = strconv.Itoa(465464)
	//test["sign"] = "sfdhdfsh"
	SignTest(da...)
}

var da = []string{"time", "suffers", "sign", "123", "data", "thteh", "domain", "http://127.0.0.1"}


func SignTest(data ...string)  {
	test := make(map[string]string)
	var ss string
	for k, v := range data {
		if k%2 == 0 {
			ss = v
		}
		test[ss] = v
	}
	fmt.Println(test)
}