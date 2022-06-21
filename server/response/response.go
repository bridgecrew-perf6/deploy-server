package response

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"oneclick/server/Signature"
	"oneclick/server/pkg"
	"strconv"
)

var sign string

func Response(c *gin.Context, httpStatus int, ret int, msg string, data gin.H) {
	m, _ := json.Marshal(&data)
	sign = Signature.SignatureMethod{}.RequestSucceeded(strconv.Itoa(ret), msg, string(m))
	c.JSON(httpStatus, gin.H{
		"ret":  ret,
		"msg":  msg,
		"sign": sign,
		"data": data,
	})
}

func Success(c *gin.Context, msg string, data gin.H) {
	Response(c, http.StatusOK, 1, msg, data)
}

func Fail(c *gin.Context, msg string, data gin.H) {
	Response(c, http.StatusOK, 0, msg, data)
}

func SignError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ReturnErrors(0, "签名错误"))
}

func ReturnErrors(ret int, msg string) (res pkg.Errors) {
	res = pkg.Errors{
		Returned: pkg.Returned{
			Ret: ret,
			Msg: msg,
		},
		Request: pkg.Request{
			Sign: sign,
		},
	}
	sign = Signature.SignatureMethod{}.RequestSucceeded(strconv.Itoa(ret), msg, "")
	res.Sign = sign
	fmt.Println(res)
	return res
}

//func ReturnErrors(ret int, msg string, content string) (res pkg.Errors) {
//	res = pkg.Errors{
//		Returned: pkg.Returned{
//			Ret: ret,
//			Msg: msg,
//		},
//		Request: pkg.Request{
//			Sign: sign,
//		},
//		Data: struct {
//			Content string `json:"content"`
//		}{
//			Content: content,
//		},
//	}
//	m, _ := json.Marshal(&res.Data)
//	sign = Signature.SignatureMethod{}.RequestSucceeded(strconv.Itoa(ret), msg, string(m))
//	res.Sign = sign
//	return res
//}
