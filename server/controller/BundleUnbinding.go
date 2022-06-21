package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oneclick/server/Signature"
	"oneclick/server/controller/methods"
	"oneclick/server/pkg"
	"runtime"
	"strconv"
)

func BundleUnbinding(c *gin.Context) {
	fmt.Println(runtime.NumGoroutine())
	var data pkg.ClientBinding
	err := c.ShouldBindJSON(&data)
	if err != nil {
		//返回错误
		methods.ErrorDesc(c, err.Error())
		return
	}
	//验签
	solution := make(map[string]string)
	solution["time"] = strconv.Itoa(data.Time)
	solution["sign"] = data.Sign
	solution["key"] = data.Key
	signature := Signature.SignatureMethod{}.KuratowskiConstraint(solution, data.Sign)
	if !signature {
		//返回签名错误
		methods.SignError(c)
		return
	}
	//_, key, _ := processor.SelectMysql(data.Key)
	//if len(key) == 0 {
	//	c.JSON(http.StatusOK, returned.ReturnErrors(0, "请求失败", "密钥不存在"))
	//	return
	//}
	//Bool := processor.Processor{}.DeleteData(data.Key)
	//if Bool {
	//	c.JSON(http.StatusOK, returned.ReturnErrors(1, "请求成功", "解除绑定成功"))
	//	return
	//}
	//c.JSON(http.StatusBadRequest, returned.ReturnErrors(0, "请求失败", "解除绑定失败"))
	//return
}
