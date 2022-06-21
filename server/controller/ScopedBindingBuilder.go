package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oneclick/server/Signature"
	"oneclick/server/controller/methods"
	"oneclick/server/pkg"
	"oneclick/server/processor"
	"oneclick/server/response"
	"runtime"
	"strconv"
)

func ScopedBindingBuilder(c *gin.Context) {
	var data pkg.ClientBinding
	err := c.ShouldBindJSON(&data)
	if err != nil {
		//返回错误
		response.Fail(c, "请求失败", gin.H{
			"content": err.Error(),
		})
		return
	}
	//验签
	solution := make(map[string]string)
	solution["time"] = strconv.Itoa(data.Time)
	solution["sign"] = data.Sign
	solution["domain"] = data.Domain
	solution["networklaundry"] = strconv.Itoa(data.Networklaundry)
	solution["key"] = data.Key
	signature := Signature.SignatureMethod{}.KuratowskiConstraint(solution, solution["sign"])
	if !signature {
		//返回签名错误
		methods.SignError(c)
		return
	}
	result := processor.Processor{}.Create()
	if result != nil {
		response.Fail(c, "请求失败", gin.H{
			"content": "密钥不存在",
		})
		return
	}
	//err := processor.Processor{}.Create(data.Key)
	//if err != nil {
	//	methods.ErrorDesc(c, "请求失败", "绑定失败，请重试")
	//	return
	//}
	//switch {
	//case !result:
	//	response.Fail(c, "请求失败",  gin.H{
	//		"content": "密钥不存在",
	//	})
	//	return
	//case len(key) == 0:
	//	methods.ErrorDescs(c, "请求失败", "密钥已绑定")
	//	return
	//case ReturnResults.Networklaundry == 1:
	//	Processor{}.insertData(key)
	//	c.JSON(200, returned.SuccessWg(1, "请求成功", utils.WireguardConfig, utils.FrpcConfig))
	//case ReturnResults.Networklaundry == 2:
	//	Processor{}.insertData(key)
	//	c.JSON(200, returned.ReturnUdp2raw(1, "请求成功", 4000, "127.0.0.1:50000", "speedbox", utils.WireguardConfig, utils.FrpcConfig))
	//}
	fmt.Println(runtime.NumGoroutine())
}
